package main

import (
	"github.com/mikebarkmin/go-artsy/internal/natural"
	"github.com/mikebarkmin/go-artsy/pkg/art"
)

func main() {
	art := natural.New(3000, 3000)
	art.SetSeed(8574979046366121729)
	single(art)
}

func single(art art.Renderable) {
	art.Render()
}

func batch(art art.Renderable) {
	for i := 0; i < 30; i++ {
		art.GenSeed()
		art.Render()

	}
}
