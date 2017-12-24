LINUX_ARCH:=amd64 arm arm64 ppc64le s390x

all: docker

.PHONY: bin
bin: build dockerfile tag qemu

.PHONY: build
build:
	cd build
	$(MAKE)

.PHONY: dockerfile
dockerfile:
	cd ./dockerfile
	$(MAKE)

.PHONY: tag
tag:
	cd ./tag
	$(MAKE)

.PHONY: qemu
qemu:
	cd ./qemu
	$(MAKE)
