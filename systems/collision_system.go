package systems

import (
	"github.com/BrianAnakPintar/ducktape/components"
	"github.com/BrianAnakPintar/ducktape/grid"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type CollisionSystem struct {
	query *donburi.Query
	grid *grid.UniformGrid // grid in struct to ensure persistence for static between updates
}

func NewCollisionSystem(w *donburi.World, g *grid.UniformGrid) *CollisionSystem {
	return &CollisionSystem{
		query: donburi.NewQuery(filter.Contains(components.Transform, components.Collider)),
		grid: g,
	}
}

func (cs *CollisionSystem) Update(world donburi.World) {
	// clear bucket full of last frame's dynamic entities to refresh
	cs.grid.ClearDynamic()

	var dynamicObjects []*donburi.Entry
	for entry := range cs.query.Iter(world) {
		c := components.Collider.Get(entry)
		if c.Static {
			continue
		}
		dynamicObjects = append(dynamicObjects, entry)

		t := components.Transform.Get(entry)
		aabb := grid.AABB{
			MinX: t.Pos.X - c.HalfWidth,
			MinY: t.Pos.Y - c.HalfHeight,
			MaxX: t.Pos.X + c.HalfWidth,
			MaxY: t.Pos.Y + c.HalfHeight,
		}
		cs.grid.InsertDynamic(int(entry.Id()), aabb)
	}
}
