package systems

import (
	"math"

	"github.com/BrianAnakPintar/ducktape/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PhysicsSystem struct {
	query *donburi.Query
}

func NewPhysicsSystem() *PhysicsSystem{
	return &PhysicsSystem{
		query: donburi.NewQuery(filter.Contains(components.Velocity, components.Transform)),
	}
}

// degToRad converts degrees to radians.
func degToRad(d float64) float64 {
    return d * (math.Pi / 180)
}

func (p *PhysicsSystem) Update(world donburi.World) {
	for entry := range p.query.Iter(world) {
		transform := components.Transform.Get(entry)
		vel := components.Velocity.Get(entry)
		transform.Pos.X += vel.PosVelocity.X
		transform.Pos.Y += vel.PosVelocity.Y

		rad := degToRad(vel.RotVelocity)
		transform.Rot = math.Mod(transform.Rot + rad, 2 * math.Pi)	
	}
}

