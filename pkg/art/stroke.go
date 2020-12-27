package art

type stroke struct {
	main  *line
	lines []*line
}

func NewStroke(l *line) *stroke {
	s := &stroke{main: l}

	var lines []*line

	for i := 0; i < 10; i++ {
		p := *l
		p.Jitter(-5, 5)
		r := RandFloat64(-5, 5)
		angle := RandFloat64(0, 360)

		p.start.TranslatePolar(r, angle)
		r = RandFloat64(-5, 5)
		angle = RandFloat64(0, 360)
		p.end.TranslatePolar(r, angle)

		p.SmoothChaikin(1)
		p.SubdivideSegments(20)
		lines = append(lines, &p)
	}
	s.lines = lines
	return s
}
