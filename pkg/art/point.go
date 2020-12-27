package art

import (
	"math"
)

type point struct {
	X float64
	Y float64
}

func (p *point) TranslatePolar(radius float64, degrees float64) {
	radians := math.Pi * degrees / 180.0
	p.X = p.X + radius*math.Cos(radians)
	p.Y = p.Y + radius*math.Sin(radians)
}

func (p *point) DistanceTo(p2 *point) float64 {
	return math.Sqrt(math.Pow(p.X-p2.X, 2) + math.Pow(p.Y-p2.Y, 2))
}

func (p *point) AngleBetween(p2 *point) float64 {
	radians := math.Atan2(p2.Y-p.Y, p2.X-p.X)
	degrees := radians * 180 / math.Pi

	if degrees < 0 {
		degrees += 360
	}

	return degrees
}

func (p *point) Copy() *point {
	point := &point{X: p.X, Y: p.Y}
	return point
}

func NewPoint(x float64, y float64) *point {
	point := &point{X: x, Y: y}
	return point
}
