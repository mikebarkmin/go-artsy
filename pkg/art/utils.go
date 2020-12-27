package art

import (
	"math/rand"
)

func RandFloat64(min float64, max float64) float64 {
	r := rand.Float64()*(max-min) + min
	return r
}

func HaltonSequence(index int, base int) float64 {
	f := 1.0
	r := 0.0

	for index > 0 {
		f /= float64(base)
		r += f * float64((index % base))
		index = int(index / base)
	}

	return r
}

func HaltonSequence23(index int) *point {
	x := HaltonSequence(index, 2)
	y := HaltonSequence(index, 3)

	return NewPoint(x, y)
}
