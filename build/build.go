package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/miekg/dxbuild"
)

const me = "/usr/bin/build"

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

	err = os.Link(me, "/bin/sh")
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
	for _, bin := range dxbuild.Archs {
		os.Remove(bin)
	}
	os.Remove(me)
	os.Remove("/usr/bin/cross-build-clean")
	os.Remove("/usr/bin/cross-build-end")
	os.Remove("/usr/bin/cross-build-start")
}

// shell runs a shell command.
func shell() error {
	var cmd *exec.Cmd

	options := append([]string{"-0", os.Args[0], "/bin/sh"}, os.Args[1:]...)

	for _, bin := range dxbuild.Archs {
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

		if err := shell(); err != nil {
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
