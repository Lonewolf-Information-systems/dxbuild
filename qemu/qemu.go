package main

import (
	"fmt"
	"log"
	"os"

	"github.com/miekg/dxbuild"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need arch")
	}
	qemu, ok := dxbuild.Archs[os.Args[1]]
	if !ok {
		log.Fatalf("Can find %s in archs map", os.Args[1])
	}
	// strip off /usr/bin/
	fmt.Println(qemu[9:])
}
