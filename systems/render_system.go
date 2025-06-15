package systems

import (
	"github.com/BrianAnakPintar/ducktape/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type RenderSystem struct {
	query *donburi.Query
}

func NewRenderSystem() *RenderSystem{
	return &RenderSystem{
		query: donburi.NewQuery(filter.Contains(components.Sprite, components.Transform)),
	}
}

// This function draws all sprite components onto the screen. (It's runned every frame)
func (a *RenderSystem) Draw(world donburi.World, screen *ebiten.Image)  {
	for entry := range components.Sprite.Iter(world) {
		spriteData := components.Sprite.Get(entry)
		transformData := components.Transform.Get(entry)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(transformData.Pos.X, transformData.Pos.Y)

		screen.DrawImage(spriteData.Image, op)
	}
}

