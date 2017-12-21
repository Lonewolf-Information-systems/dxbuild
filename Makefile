dxbuild: dxbuild.go
	go build -ldflags "-w -s" dxbuild.go

.PHONY: bin
bin: dxbuild
	mkdir -p bin
	cp $$(which qemu-arm-static) bin/qemu-arm-static
	cp dxbuild bin
	ln -sf dxbuild bin/cross-build-end
	ln -sf dxbuild bin/cross-build-start
	ln -sf dxbuild bin/cross-build-clean
	ln -sf sh bin/sh.real

.PHONY: docker
docker: bin
	docker build -t arm32v6/builder:latest .
