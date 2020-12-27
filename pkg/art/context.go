package art

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/fogleman/gg"
)

type Context struct {
	seed   int64
	perlin *perlin.Perlin
	gg.Context
}

func (c *Context) Deg2Rad(degrees float64) float64 {
	return gg.Radians(degrees)
}

func (c *Context) Rad2Deg(radians float64) float64 {
	return gg.Degrees(radians)
}

func (c *Context) DrawBackground() {
	c.DrawRectangle(0, 0, float64(c.Width()), float64(c.Height()))
}

func (c *Context) Render(name string) error {
	fmt.Printf("Seed: %d\n", c.seed)
	return c.SavePNG(fmt.Sprintf("out/%s_%d.png", name, c.seed))
}

func (c *Context) DrawLine(l *line) {
	c.MoveTo(l.start)
	for _, p := range l.inbetween {
		c.LineTo(p)
	}
	c.LineTo(l.end)
}

func (c *Context) DrawStroke(s *stroke) {
	for _, l := range s.lines {
		c.DrawLineAsPoints(l, 1)
		c.Stroke()
		c.Fill()
	}
	c.DrawLine(s.main)
	c.Stroke()
}

func (c *Context) DrawLineAsPoints(l *line, radius float64) {
	c.DrawPoint(l.start, radius)
	for _, p := range l.inbetween {
		c.DrawPoint(p, radius)
	}
	c.DrawPoint(l.end, radius)
}

func (c *Context) DrawPoint(p *point, radius float64) {
	c.Context.DrawPoint(p.X, p.Y, radius)
}

func (c *Context) MoveTo(p *point) {
	c.Context.MoveTo(p.X, p.Y)
}

func (c *Context) LineTo(p *point) {
	c.Context.LineTo(p.X, p.Y)
}

func (c *Context) SetPixel(p *point) {
	c.Context.SetPixel(int(p.X), int(p.Y))
}

func (c *Context) GetSeed() int64 {
	return c.seed
}

func (c *Context) SetSeed(seed int64) {
	c.seed = seed
	rand.Seed(seed)
}

func (c *Context) GenSeed() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	seed := int64(binary.LittleEndian.Uint64(b[:]))
	// only allow positive seeds. This is better for filenames.
	if seed < 0 {
		seed *= -1
	}
	c.seed = seed
	rand.Seed(seed)
}

func NewContext(width int, height int) *Context {
	c := gg.NewContext(width, height)
	a := &Context{0, nil, *c}
	a.GenSeed()

	p := perlin.NewPerlin(2, 2, 3, a.seed)
	a.perlin = p

	return a
}
