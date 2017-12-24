package main

import (
	"html/template"
	"log"
	"os"

	"github.com/miekg/dxbuild"
)

// dockerFile defines our dockerfile
type dockerFile struct {
	Image string
	Qemu  string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need arch")
	}
	debian, ok := dxbuild.DebianImages[os.Args[1]]
	if !ok {
		log.Fatalf("Can find %s in image map", os.Args[1])
	}
	qemu, ok := dxbuild.Archs[os.Args[1]]
	if !ok {
		log.Fatalf("Can find %s in qemu map", os.Args[1])
	}

	t := template.Must(template.New("Dockerfile.tmpl").ParseFiles("Dockerfile.tmpl"))
	df := dockerFile{Image: debian, Qemu: qemu[9:]} // strip off /usr/bin/
	if err := t.Execute(os.Stdout, df); err != nil {
		log.Fatal(err)
	}
}
