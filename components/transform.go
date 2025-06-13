package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type Transform struct {
	Pos math.Vec2
	Rot float64
}

var Transforms = donburi.NewComponentType[Transform]()

