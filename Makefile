LINUX_ARCH:=amd64 arm arm64 ppc64le s390x

all: docker

.PHONY: docker
docker: bin
	docker build -t arm32v6/builder:latest .

.PHONY: bin
bin: dxbuild dockerfile tag qemu

.PHONY: dxbuild
dxbuild:
	cd dxbuild
	$(MAKE)

.PHONY: dockerfile
dockerfile:
	cd dockerfile
	$(MAKE)

.PHONY: tag
tag:
	cd tag
	$(MAKE)

.PHONY: qemu
qemu:
	cd qemu
	$(MAKE)
