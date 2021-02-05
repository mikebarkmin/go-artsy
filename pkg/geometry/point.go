package geometry

type point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *point {
	p := &point{x, y}
	return p
}
