package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type TransformData struct {
	Pos math.Vec2
	Rot float64
}

var Transform = donburi.NewComponentType[TransformData]()

