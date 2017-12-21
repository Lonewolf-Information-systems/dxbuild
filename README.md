# dxbuild

`dxbuild` allows you to build non-amd64 containers on amd64, for reaons unknonwn to me this is hard
to do in Docker land.

This is not my idea:

See [this blog
post](https://resin.io/blog/building-arm-containers-on-any-x86-machine-even-dockerhub/), I used the
code [here](https://github.com/resin-io-projects/armv7hf-debian-qemu) and forked it to this and
cleaned it up a bit and made it support all k8s architecture.

## Usage

You'll need two docker containers (multistage builds sadly don't work). One to install the various
binaries and then another where you actually use them.

This example is *just* for arm.

Dockerfile "a":
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
cp dxbuild bin # this is a amd64 binary
~~~

`make bin` will do this for you.

