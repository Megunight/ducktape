package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type Velocity struct {
	PosVelocity math.Vec2
	RotVelocity float64
}

var Velocities = donburi.NewComponentType[Velocity]()
