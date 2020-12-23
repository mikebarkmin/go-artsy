package art

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"

	"github.com/fogleman/gg"
)

type Context struct {
	seed int64
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
	return c.SavePNG(fmt.Sprintf("out/%s.png", name))
}

func (c *Context) GetSeed() int64 {
	return c.seed
}

func (c *Context) SetSeed(seed int64) {
	c.seed = seed
	rand.Seed(seed)
}

func (c *Context) RandFloat64(min float64, max float64) float64 {
	r := rand.Float64()*(max-min) + min
	return r
}

func (c *Context) GenSeed() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	seed := int64(binary.LittleEndian.Uint64(b[:]))
	c.seed = seed
}

func NewContext(width int, height int) *Context {
	c := gg.NewContext(width, height)
	a := &Context{0, *c}
	a.GenSeed()

	return a
}
