package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

const dxbuild = "/usr/bin/dxbuild"

func crossBuildStart() {
	if _, err := os.Stat("/bin/sh.real"); os.IsNotExist(err) {
		err = os.Link("/bin/sh", "/bin/sh.real")
		if err != nil {
			log.Fatal(err)
		}
	}

	err := os.Remove("/bin/sh")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Link(dxbuild, "/bin/sh")
	if err != nil {
		log.Fatal(err)
	}
}

func crossBuildEnd() {
	err := os.Remove("/bin/sh")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Link("/bin/sh.real", "/bin/sh")
	if err != nil {
		log.Fatal(err)
	}

}

func crossBuildClean() {
	for _, bin := range archs {
		os.Remove(bin)
	}
	os.Remove(dxbuild)
	os.Remove("/usr/bin/cross-build-clean")
	os.Remove("/usr/bin/cross-build-end")
	os.Remove("/usr/bin/cross-build-start")
}

// If we find any of these we will use them.
var archs = map[string]string{
	"amd64":   "/usr/bin/qemu-x86_64-static",
	"arm":     "/usr/bin/qemu-arm-static",
	"arm64":   "/usr/bin/qemu-aarch64-static",
	"ppc64le": "/usr/bin/qemu-ppc64le-static",
	"s390x":   "/usr/bin/qemu-s390x-static",
}

func runShell() error {
	var cmd *exec.Cmd

	options := append([]string{"-0", os.Args[0], "/bin/sh"}, os.Args[1:]...)

	for _, bin := range archs {
		if _, err := os.Stat(bin); err == nil {
			cmd = exec.Command(bin, append(options)...)
			break
		}
	}
	if cmd == nil {
		log.Fatal("no qemu-*-static found in /usr/bin")
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	switch os.Args[0] {
	case "cross-build-start":
		crossBuildStart()
	case "cross-build-end":
		crossBuildEnd()
	case "cross-build-clean":
		crossBuildClean()
	default:
		code := 0
		crossBuildEnd()

		if err := runShell(); err != nil {
			code = 1
			if exiterr, ok := err.(*exec.ExitError); ok {
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					code = status.ExitStatus()
				}
			}
		}

		crossBuildStart()

		os.Exit(code)
	}
}
