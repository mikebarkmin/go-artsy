package xmastree

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

type xmastree struct {
	width   float64
	height  float64
	context gg.Context
}

func (xt *xmastree) tree(startx float64, starty float64, length float64, rotation float64, lineWidth float64) {
	c := xt.context
	endx := startx + length*math.Cos(gg.Radians(rotation))
	endy := starty + length*math.Sin(gg.Radians(rotation))
	c.DrawLine(startx, starty, endx, endy)
	c.SetLineWidth(lineWidth)
	// make gradient darker
	gradientColor1 := math.Max((float64(xt.height)-starty)/float64(xt.height)-0.2, 0)
	gradientColor2 := math.Max((float64(xt.width)-startx)/float64(xt.width)-0.2, 0)
	c.SetRGB(gradientColor1, 0.5, gradientColor2)
	c.Stroke()

	if length > 5 {
		newLineWidth := math.Max(lineWidth-0.5, 2)
		newLength := length / 1.75
		newLengthCenter := length - 2
		xt.tree(endx, endy, newLength, 50+rotation, newLineWidth)
		xt.tree(endx, endy, newLength, -50+rotation, newLineWidth)
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

func (xt *xmastree) Render(path string) {
	fmt.Println("Rendering")
	c := xt.context
	c.Clear()

	xt.background()
	xt.stars(50)
	xt.tree(float64(xt.width)/2.0, float64(xt.height), 80, -90, 6)
	xt.stars(20)

	fmt.Println("Saving")
	c.SavePNG(fmt.Sprintf("%s.png", path))
}

func New(width float64, height float64) (*xmastree, error) {
	if width < 100 {
		return nil, fmt.Errorf("width of %d is under the minimum of 100", width)
	}
	if height < 100 {
		return nil, fmt.Errorf("height of %d is under the minimum of 100", height)
	}
	c := gg.NewContext(int(width), int(height))
	x := &xmastree{width: width, height: height, context: *c}
	return x, nil
}
