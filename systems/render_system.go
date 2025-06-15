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
		query: donburi.NewQuery(filter.Contains(components.Sprite)),
	}
}

func (a *RenderSystem) Draw(world donburi.World, screen *ebiten.Image)  {
	for entry := range components.Sprite.Iter(world) {
		spriteData := components.Sprite.Get(entry)
		screen.DrawImage(spriteData.Image, nil)
	}
}

