package looplooploop

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mikebarkmin/go-artsy/pkg/art"
)

type looplooploop struct {
	art.Context
}

func (l *looplooploop) Render() error {
	l.SetRGB(1, 1, 1)
	l.DrawBackground()
	l.Fill()

	width := l.Width64()
	height := l.Height64()

	margin := 150.0
	innerWidth := width - 2*margin
	innerHeight := height - 2*margin

	hueFix := 180.0
	huePos := 120.0
	hueRand := 10.0
	// First Loop
	for x := 0; x < 200; x++ {
		// Second Loop
		for y := 0; y < 200; y++ {
			index := x + y*l.Width()
			// Third hidden Loop
			p := art.HaltonSequence23(index)
			hueR := art.RandFloat64(-hueRand, hueRand)
			satR := art.RandFloat64(-0.2, 0.2)
			c := colorful.Hsv(hueFix+p.Y*huePos+hueR, p.X+0.9*satR, 0.9).Clamped()

			p.X *= innerWidth
			p.X += margin
			p.Y *= innerHeight
			p.Y += margin

			radius := art.RandFloat64(10, 40)
			rotation := art.RandFloat64(40, 50)

			l.DrawRegularPolygon(4, p.X, p.Y, radius, rotation)
			l.SetColor(c)
			l.SetLineWidth(0)
			l.Fill()
		}
	}

	return l.Context.Render("looplooploop")
}

func New(width int, height int) *looplooploop {
	c := art.NewContext(width, height)
	l := &looplooploop{*c}
	return l
}
