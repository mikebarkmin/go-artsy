package art

import "github.com/fogleman/gg"

type Art interface {
	Draw() gg.Context
}
