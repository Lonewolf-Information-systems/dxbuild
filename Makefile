.PHONY: bin
bin: dxbuild
	mkdir -p bin
	cp $$(which qemu-arm-static) bin/qemu-arm-static
	cp dxbuild bin
	ln -sf dxbuild bin/cross-build-end
	ln -sf dxbuild bin/cross-build-start
	ln -sf sh bin/sh.real


dxbuild:
	go build -ldflags "-w -s" dxbuild.go

.PHONY: docker
docker:
	docker build -t arm32v6/builder:latest .
