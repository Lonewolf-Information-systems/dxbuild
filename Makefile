

PHONY: bin
bin:
	mkdir -f bin
	cp $$(which qemu-arm-static) bin/qemu-arm-static
	ln -sf dxbuild bin/cross-build-end
	ln -sf dxbuild bin/cross-build-start
	ln -sf sh bin/sh.real


dxbuild:
	go build -ldflags "-w -s" dxbuild.go
