# dxbuild

`dxbuild` allows you to build non-amd64 containers on amd64.

**This is not my idea**

See [this blog
post](https://resin.io/blog/building-arm-containers-on-any-x86-machine-even-dockerhub/), I used the
code [here](https://github.com/resin-io-projects/armv7hf-debian-qemu) and forked it to this and
cleaned it up a bit and made it possible support all k8s architectures.

I cleaned it up a bit and added some documentation and make it work (my qemu didn't have some
options).

This has been tested on Debian. You'll also need [qemu](https://wiki.debian.org/QemuUserEmulation)
installed, but this boils down on apt-get installing some things.

## Usage

~~~
% make docker
~~~
And then create a *second* Dockerfile with the non amd64 image you want to build, i.e.
~~~
FROM arm32v6/builder:latest

RUN [ "cross-build-start" ]
RUN apk --update add stunnel
RUN [ "cross-build-end" ]
RUN [ "cross-build-clean" ]
~~~

More below.

You'll need two docker containers (multistage builds sadly don't work). One to install the various
binaries and then another where you actually use them.

This example is *just* for ARM.

Dockerfile:
~~~
FROM arm32v6/alpine:latest
COPY bin/ usr/bin/
~~~

Where `bin` needs to be populated with:

~~~ sh
cp $(which qemu-arm-static) bin/qemu-arm-static
ln -sf dxbuild bin/cross-build-end
ln -sf dxbuild bin/cross-build-start
ln -sf sh bin/sh.real
cp dxbuild bin # this is an amd64 binary
~~~

`make bin` will do this for you. And `make docker` will create a arm32v6/builder:latest image.
*That* image can now be used to setup a proper arm32v6 image, i.e., Dockerfile:

~~~
FROM arm32v6/builder:latest

RUN [ "cross-build-start" ]

RUN apk --update add stunnel

RUN [ "cross-build-end" ]
RUN [ "cross-build-clean" ]
~~~

And just build that like normal, `dxbuild` does remove itself from the image (cross-build-clean),
but layers.

If you execute the above it will looks like this:

~~~ sh
% docker build -t miekg/stunnel
Sending build context to Docker daemon  2.048kB
Step 1/5 : FROM arm32v6/builder:latest
 ---> 8f58ca070d3f
Step 2/5 : RUN cross-build-start
 ---> Running in 3a3ffec4fb06
 ---> f09bb9af2e24
Removing intermediate container 3a3ffec4fb06
Step 3/5 : RUN apk --update add stunnel
 ---> Running in 659222a958bc
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/main/armhf/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.6/community/armhf/APKINDEX.tar.gz
(1/1) Installing stunnel (5.38-r1)
Executing stunnel-5.38-r1.pre-install
Executing busybox-1.26.2-r5.trigger
OK: 3 MiB in 12 packages
 ---> 1c5f121c0e24
Removing intermediate container 659222a958bc
Step 4/5 : RUN cross-build-end
 ---> Running in a57e8244c2f4
 ---> 25e99b362306
Removing intermediate container a57e8244c2f4
Step 5/5 : RUN cross-build-clean
 ---> Running in 0b79feb0976e
 ---> f4ad80f78f2c
Removing intermediate container 0b79feb0976e
Successfully built f4ad80f78f2c
Successfully tagged miekg/stunnel:latest
~~~

Other architectures are supported, `dxbuild` will check for:

~~~ golang
var archs = map[string]string{
	"amd64":   "/usr/bin/qemu-x86_64-static",
	"arm":     "/usr/bin/qemu-arm-static",
	"arm64":   "/usr/bin/qemu-aarch64-static",
	"ppc64le": "/usr/bin/qemu-ppc64le-static",
	"s390x":   "/usr/bin/qemu-s390x-static",
}
~~~
And will execute the first one found and this, of course, depends on what you copied into bin/.
