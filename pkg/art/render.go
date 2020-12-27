package art

type Renderable interface {
	Render() error
	GenSeed()
	SetSeed(seed int64)
}
