package systems

import (
	"github.com/BrianAnakPintar/ducktape/components"
	c "github.com/BrianAnakPintar/ducktape/constants"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type GravitySystem struct {
	query *donburi.Query
}

func NewGravitySystem() *GravitySystem{
	return &GravitySystem{
		query: donburi.NewQuery(filter.Contains(components.Velocity, components.Transform)),
	}
}

func (p *GravitySystem) Update(world donburi.World) {
	for entry := range p.query.Iter(world) {
		vel := components.Velocity.Get(entry)
		vel.PosVelocity.Y += c.PlayerGravity
	}
}

