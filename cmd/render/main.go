package main

import (
	"github.com/mikebarkmin/go-artsy/pkg/xmastree"
)

func main() {
	art, err := xmastree.New(1500, 1500)
	if err != nil {
		panic(err)
	}

	art.Render("out/xmastree")
}
