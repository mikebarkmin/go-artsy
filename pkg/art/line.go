package art

type line struct {
	start     *point
	end       *point
	inbetween []*point
}

func (l *line) Segmentize(number int) {
	inbetween := []*point{}
	length := l.GetLuftline()
	segmentLength := length / float64(number)

	if length <= segmentLength {
		return
	}

	angle := l.GetAngle()

	currentPoint := l.start
	for i := 0; i < number-1; i++ {
		n := *currentPoint
		n.TranslatePolar(segmentLength, angle)
		currentPoint = &n
		inbetween = append(inbetween, currentPoint)
	}
	l.inbetween = inbetween
}

func (l *line) SubdivideSegments(number int) {
	points := []*point{l.start}
	points = append(points, l.inbetween...)
	points = append(points, l.end)

	newInbetween := []*point{}

	var nextPoint point
	for i := 0; i+1 < len(points); i++ {
		currentPoint := *points[i]
		nextPoint = *points[i+1]

		tmpLine := NewLine(&currentPoint, &nextPoint)
		tmpLine.Segmentize(number)

		newInbetween = append(newInbetween, &currentPoint)
		newInbetween = append(newInbetween, tmpLine.inbetween...)
	}

	newInbetween = append(newInbetween, &nextPoint)

	l.inbetween = newInbetween
}

func (l *line) GetAngle() float64 {
	return l.start.AngleBetween(l.end)
}

func (l *line) GetLuftline() float64 {
	return l.start.DistanceTo(l.end)
}

func (l *line) JitterY(min float64, max float64) {
	for i := range l.inbetween {
		r := RandFloat64(min, max)
		l.inbetween[i].Y += r
	}
}

func (l *line) JitterX(min float64, max float64) {
	for i := range l.inbetween {
		r := RandFloat64(min, max)
		l.inbetween[i].X += r
	}
}

func (l *line) Jitter(minRadius float64, maxRadius float64) {
	for i := range l.inbetween {
		r := RandFloat64(minRadius, maxRadius)
		angle := RandFloat64(0, 360)
		l.inbetween[i].TranslatePolar(r, angle)
	}
}

func (l *line) SmoothNeighbor(iterations int) {

}

func (l *line) SmoothChaikin(iterations int) {

	var inner func(points []*point, iterations int) []*point
	inner = func(points []*point, iterations int) []*point {

		if iterations == 0 {
			return points
		}

		l := len(points)

		var smoothPoints []*point

		for i := 0; i < l-1; i++ {
			p := points[i]
			nextPoint := points[(i+1)%l]
			nextX := nextPoint.X
			nextY := nextPoint.Y
			x := p.X
			y := p.Y
			smoothX1 := 0.75*x + 0.25*nextX
			smoothY1 := 0.75*y + 0.25*nextY
			smoothX2 := 0.25*x + 0.75*nextX
			smoothY2 := 0.25*y + 0.75*nextY
			point1 := NewPoint(smoothX1, smoothY1)
			point2 := NewPoint(smoothX2, smoothY2)
			smoothPoints = append(smoothPoints, point1, point2)
		}

		if iterations == 1 {
			return smoothPoints
		}

		return inner(smoothPoints, iterations-1)
	}

	points := append(l.inbetween, l.end)
	points = append([]*point{l.start}, points...)
	l.inbetween = inner(points, iterations)
}

func (l *line) Copy() *line {
	n := new(line)
	*n = *l

	for _, p := range l.inbetween {
		q := new(point)
		*q = *p
		n.inbetween = append(n.inbetween, q)
	}

	return n
}

func NewLine(start *point, end *point) *line {
	line := &line{start: start, end: end}
	return line
}
