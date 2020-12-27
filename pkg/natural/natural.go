package natural

import (
	"github.com/mikebarkmin/go-artsy/pkg/art"
)

type natural struct {
	art.Context
}

func (n *natural) Render() error {

	margin := 300.0
	gap := 220.0
	innerGap := 150.0

	n.SetRGB(0.86, 0.86, 0.86)
	n.DrawBackground()
	n.Fill()
	n.Stroke()

	for i := 0; i < 120000; i++ {
		p := art.HaltonSequence23(i)
		width := float64(n.Width()) - 2*innerGap - 2*margin
		height := float64(n.Height()) - 2*innerGap
		p.X = p.X*width + innerGap + margin
		p.Y = p.Y*height + innerGap

		r := art.RandFloat64(0.0, 0.1)
		n.SetRGB(0.72+r, 0.72+r, 0.72+r)
		n.DrawPoint(p, 2+art.RandFloat64(-0.5, 0.5))
		n.Fill()
		n.Stroke()
	}

	n.DrawRectangle(margin*2-innerGap, margin-innerGap, float64(n.Width())-4*margin+2*innerGap, float64(n.Height())-2*margin+2*innerGap)
	n.SetRGB(0, 0, 0)
	n.SetLineWidth(8)
	n.Stroke()

	for i := 0.0; i < 12; i++ {
		p1 := art.NewPoint(margin, i*gap+margin)
		p2 := art.NewPoint(float64(n.Width())-margin, i*gap+margin)

		l := art.NewLine(p1, p2)
		l.Segmentize(3)
		l.Jitter(40, 50)
		l.SmoothChaikin(2)

		s := art.NewStroke(l)
		n.SetRGB(1, 0, 0)
		n.SetLineWidth(2)
		n.DrawStroke(s)
	}

	return n.Context.Render("natural")
}

func New(width int, height int) *natural {
	c := art.NewContext(width, height)
	n := &natural{*c}
	return n
}
