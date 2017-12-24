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
	build, ok := dxbuild.DebianBuildImages[os.Args[1]]
	if !ok {
		log.Fatalf("Can find %s in image map", os.Args[1])
	}
	fmt.Println(build)
}
