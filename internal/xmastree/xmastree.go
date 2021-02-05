package xmastree

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/mikebarkmin/go-artsy/pkg/art"
)

type xmastree struct {
	width   float64
	height  float64
	context art.Context
}

func (xt *xmastree) tree(startx float64, starty float64, length float64, rotation float64, lineWidth float64) {
	c := xt.context
	endx := startx + length*math.Cos(gg.Radians(rotation))
	endy := starty + length*math.Sin(gg.Radians(rotation))

	p1 := art.NewPoint(startx, starty)
	p2 := art.NewPoint(endx, endy)
	l := art.NewLine(p1, p2)
	c.DrawLine(l)
	c.SetLineWidth(lineWidth)
	// make gradient darker
	gradientColor1 := math.Max((float64(xt.height)-starty)/float64(xt.height)-(0.4*rand.Float64()+0.1), 0)
	gradientColor2 := math.Max((float64(xt.width)-startx)/float64(xt.width)-(0.4*rand.Float64()+0.1), 0)
	c.SetRGB(gradientColor1, 0.5, gradientColor2)
	c.Stroke()

	if length > 5 {
		// newLineWidth := math.Max(lineWidth-0.2, 2) // no variance
		newLineWidth := math.Max(lineWidth-(rand.Float64()+0.2), 2)
		newLength := length / 1.75
		newLengthCenter := length - 2.5
		// angle1 := 50 // no variance
		// angle2 := 50
		angle1 := rand.Float64()*14 - 7 + 50
		angle2 := rand.Float64()*14 - 7 + 50
		xt.tree(endx, endy, newLength, angle1+rotation, newLineWidth)
		xt.tree(endx, endy, newLength, -angle2+rotation, newLineWidth)
		xt.tree(endx, endy, newLengthCenter, 0+rotation, newLineWidth)
	}
}

func (xt *xmastree) stars(number int) {
	c := xt.context
	for i := 0; i < number; i++ {
		x := rand.Float64() * float64(xt.width)
		y := rand.Float64() * float64(xt.height)
		r := rand.Float64()*6 + 1
		c.DrawCircle(x, y, r)
		c.SetRGB(1, 1, 1)
		c.Fill()
	}
}

func (xt *xmastree) background() {
	c := xt.context
	c.DrawRectangle(0, 0, xt.width, xt.height)
	grad := gg.NewLinearGradient(xt.width/2.0, 0, xt.width/2.0, xt.height)
	grad.AddColorStop(0, color.RGBA{0, 0, 0, 255})
	grad.AddColorStop(0.5, color.RGBA{0, 0, 10, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 40, 255})
	c.SetFillStyle(grad)
	c.Fill()
}

func (xt *xmastree) Render() error {
	fmt.Println("Rendering")
	c := xt.context
	c.Clear()

	// xt.background()
	// xt.stars(50)
	xt.tree(float64(xt.width)/2.0, float64(xt.height), 80, -90, 8)
	// xt.stars(20)

	fmt.Println("Saving")
	return c.Render("xmastree")
}

func New(width int, height int) *xmastree {
	c := art.NewContext(width, height)
	x := &xmastree{width: float64(c.Width()), height: float64(c.Height()), context: *c}
	return x
}
