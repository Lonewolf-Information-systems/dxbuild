LINUX_ARCH:=amd64 arm arm64 ppc64le s390x

all: dxbuild dxdocker


.PHONY: bin
bin: dxbuild
	mkdir -p bin
	cp $$(which qemu-arm-static) bin/qemu-arm-static
	cp dxbuild bin

.PHONY: docker
docker: bin
	docker build -t arm32v6/builder:latest .

.PHONY: dxbuild
dxbuild:
	cd dxbuild
	$(MAKE)

.PHONY: dxdocker
dxdocker:
	cd dxdocker
	$(MAKE)
