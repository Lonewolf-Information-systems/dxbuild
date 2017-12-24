# dxbuild

`dxbuild` allows you to build non-amd64 containers on amd64. It currently uses small Debian images.

Working Golang compile chain is expected.

**This is not my idea**

See [this blog
post](https://resin.io/blog/building-arm-containers-on-any-x86-machine-even-dockerhub/), I used the
code [here](https://github.com/resin-io-projects/armv7hf-debian-qemu) and forked it to this and
cleaned it up a bit and made it possible support all k8s architectures.

This has been tested on Debian. You'll also need [qemu](https://wiki.debian.org/QemuUserEmulation)
installed, but this boils down on apt-get installing some things.

Each new image will be tagged with `builder` in it, except for amd64:

* amd64 -> debian:stable-slim
* arm -> arm32v7/debian-builder:stable-slim
* arm64 -> arm64v8/debian-builder:stable-slim
* ppc64le -> ppc64le/debian-builder:stable-slim
* s390x -> s390x/debian-builder:stable-slim

## Usage

~~~
% make        # build go binaries
% make docker # build all builder docker containers
~~~

We don't upload anything to the docker hub as we want to keep these local and just inherit from.

I.e. this is how you can use them:

~~~
FROM arm32v7/debian-builder:stable-slim

RUN [ "cross-build-start" ]

RUN apt-get update && apt-get install -y stunnel

RUN [ "cross-build-end" ]
RUN [ "cross-build-clean" ]
~~~
