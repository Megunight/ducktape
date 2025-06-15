package archetypes

import (
	"github.com/BrianAnakPintar/ducktape/assets"
	"github.com/BrianAnakPintar/ducktape/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

func NewPlayer(w donburi.World, pos math.Vec2) {
	playerEntity := w.Create(
		components.Player,
		components.Transform,
		components.Velocity,
		components.Sprite)

	entry := w.Entry(playerEntity)
	components.Transform.SetValue(entry, components.TransformData{Pos: pos, Rot: 0})
	components.Sprite.SetValue(entry, components.SpriteData{Image: assets.PlayerAsset0})
}
