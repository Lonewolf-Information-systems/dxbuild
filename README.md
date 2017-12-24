# dxbuild

`dxbuild` allows you to build non-amd64 containers on amd64. It currently uses small Debian images.

**This is not my idea**

See [this blog
post](https://resin.io/blog/building-arm-containers-on-any-x86-machine-even-dockerhub/), I used the
code [here](https://github.com/resin-io-projects/armv7hf-debian-qemu) and forked it to this and
cleaned it up a bit and made it possible support all k8s architectures.

This has been tested on Debian. You'll also need [qemu](https://wiki.debian.org/QemuUserEmulation)
installed, but this boils down on apt-get installing some things.

Each new image will be tagged with `builder` in it:

* arm32v7/debian-builder:stable-slim, for arm32v7
* arm64v8/debian-builder:stable-slim, for arm64v8

## Usage

~~~
% make docker
~~~

This will create build images that you can use in your Dockerfiles.

Create a *second* Dockerfile with the non amd64 image you want to build, i.e, here we use
the (previously created) arm32v7/debian-*builder*:stable-slim images.
~~~
FROM arm32v7/debian-builder:stable-slim

RUN [ "cross-build-start" ]
RUN apt-get install stunnel
RUN [ "cross-build-end" ]
RUN [ "cross-build-clean" ]
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
