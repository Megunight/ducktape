package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type VelocityData struct {
	PosVelocity math.Vec2
	RotVelocity float64
}

var Velocity = donburi.NewComponentType[VelocityData]()
