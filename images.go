package dxbuild

// Archs defines all the architectures we care about and have a qemu-static for.
var Archs = map[string]string{
	"amd64":   "/usr/bin/qemu-x86_64-static",
	"arm":     "/usr/bin/qemu-arm-static",
	"arm64":   "/usr/bin/qemu-aarch64-static",
	"ppc64le": "/usr/bin/qemu-ppc64le-static",
	"s390x":   "/usr/bin/qemu-s390x-static",
}

// DebianImages is a map with all the architectures we have Debian images for.
var DebianImages = map[string]string{
	"amd64":   "debian:stable-slim",
	"arm":     "arm32v7/debian:stable-slim",
	"arm64":   "arm64v8/debian:stable-slim",
	"ppc64le": "ppc64le/debian:stable-slim",
	"s390x":   "s390x/debian:stable-slim",
}

// DebianBuildImages is a map with all the architectures we generate builder images for.
var DebianBuildImages = map[string]string{
	"amd64":   "debian:stable-slim",
	"arm":     "arm32v7/debian-builder:stable-slim",
	"arm64":   "arm64v8/debian-builder:stable-slim",
	"ppc64le": "ppc64le/debian-builder:stable-slim",
	"s390x":   "s390x/debian-builder:stable-slim",
}
