package main

import (
	"github.com/mikebarkmin/go-artsy/pkg/schotter"
)

func main() {
	art := schotter.New(1500, 1500)
	art.Render()
}
