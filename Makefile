ARCH:=arm arm64 ppc64le s390x # amd64 is missing here, because we don't need it.

.PHONY: bin
bin:
	@ ( cd build; make build )
	@ ( cd dockerfile; make dockerfile )
	@ ( cd qemu; make qemu )
	@ ( cd tag; make tag )

.PHONY: docker
docker:
	@ for arch in $(ARCH); do \
	    make $$arch.docker; \
	done

# Create tmp dir, copy dockerfile into, copy build, copy qemu
%.docker:
	@ $(eval TEMP := $(shell mktemp -d))
	@ $(eval TAG := $(shell tag/tag $*))
	@ cp build/build $(TEMP)
	@ cp /usr/bin/$(shell qemu/qemu $*) $(TEMP)
	@ ./dockerfile/dockerfile $* > $(TEMP)/Dockerfile
	( cd $(TEMP); docker build -t $(TAG) . )
	@ rm -rf $(TEMP)
