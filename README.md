# dxbuild

`dxbuild` allows you to build non-amd64 containers on amd64. It currently using small Debian images.

**This is not my idea**

See [this blog
post](https://resin.io/blog/building-arm-containers-on-any-x86-machine-even-dockerhub/), I used the
code [here](https://github.com/resin-io-projects/armv7hf-debian-qemu) and forked it to this and
cleaned it up a bit and made it possible support all k8s architectures.

This has been tested on Debian. You'll also need [qemu](https://wiki.debian.org/QemuUserEmulation)
installed, but this boils down on apt-get installing some things.

## Usage

~~~
% make docker
~~~

This will create build images that you can use in your Dockerfiles.

Create a *second* Dockerfile with the non amd64 image you want to build, i.e.
~~~
FROM arm32v7/debian-builder:stable-slim

RUN [ "cross-build-start" ]
RUN apt-get install stunnel
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
