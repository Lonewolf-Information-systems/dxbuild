ARCH:=amd64 arm arm64 ppc64le s390x

.PHONY: bin
bin:
	@ ( cd build; make build )
	@ ( cd dockerfile; make dockerfile )
	@ ( cd qemu; make qemu )
	@ ( cd tag; make tag )

# Create tmp dir, copy dockerfile into, copy build, copy qemu
.PHONY: docker
docker:
	$(eval TEMP := $(shell mktemp -d))
	$(eval TAG := $(shell tag/tag arm))
	echo $(TEMP)
	cp build/build $(TEMP)
	cp /usr/bin/$(shell qemu/qemu arm) $(TEMP)
	./dockerfile/dockerfile arm > $(TEMP)/Dockerfile
	( cd $(TEMP); docker build -t $(TAG) . )
	rm -rf $(TEMP)
