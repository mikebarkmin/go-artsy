package schotter

import (
	"github.com/mikebarkmin/go-artsy/pkg/art"
)

type schotter struct {
	art.Context
}

func (s *schotter) Render() error {
	s.DrawBackground()
	s.SetRGB255(255, 255, 255)
	s.Fill()

	columns := 12.0
	size := float64(s.Width()) / columns
	rows := float64(s.Height()) / size
	padding := 1.0
	// account for padding
	columns -= padding * 2
	rows -= padding * 2
	dampen := 0.45
	randsum := 1.0
	randstep := 0.65

	for r := 0.0; r < rows; r++ {
		randsum += (r * randstep)
		for c := 0.0; c < columns; c++ {
			randval := art.RandFloat64(-randsum, randsum)
			x := float64(c) * size
			y := float64(r) * size
			x = padding*size + x + randval*dampen
			y = padding*size + y + randval*dampen
			s.Push()
			s.Translate(x, y)
			s.Rotate(s.Deg2Rad(randval))
			s.DrawRectangle(0, 0, size, size)
			s.SetLineWidth(4)
			s.SetRGB255(0, 0, 0)
			s.Stroke()
			s.Pop()
		}
	}
	return s.Context.Render("schotter")
}

func New(width int, height int) *schotter {
	c := art.NewContext(width, height)
	s := &schotter{*c}
	return s
}
