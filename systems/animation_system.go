package systems

import (
	"github.com/BrianAnakPintar/ducktape/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type AnimationSystem struct {
	query *donburi.Query
}

func NewAnimationSystem() *AnimationSystem {
	return &AnimationSystem{
		query: donburi.NewQuery(filter.Contains(components.Sprite)),
	}
}

func (a *AnimationSystem) Update(world donburi.World) {
	// TODO: Implement
}

