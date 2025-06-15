package scenes

import (
	"github.com/BrianAnakPintar/ducktape/assets"
	"github.com/BrianAnakPintar/ducktape/components"
	c "github.com/BrianAnakPintar/ducktape/constants"
	"github.com/BrianAnakPintar/ducktape/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type TestLevelScene struct {
	NumEnemies int
	text  string
	world donburi.World

	animSystem systems.AnimationSystem
	renderSystem systems.RenderSystem
}

func (t *TestLevelScene) GetName() string {
	return "TestLevel"
}

func (t *TestLevelScene) Update() {
	t.animSystem.Update(t.world)
}

func (t *TestLevelScene) Render(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, t.text)
	t.renderSystem.Draw(t.world, screen)
}

func (t *TestLevelScene) HandleInput() {
	if ebiten.IsKeyPressed(c.SkipCutsceneKey) {	
		t.text = "Brian is goated"
	}
}

func (t *TestLevelScene) OnEnterScene() {
	playerEntity := t.world.Create(components.Transform, components.Player, components.Velocity, components.Sprite)
	entry := t.world.Entry(playerEntity)

	components.Transform.SetValue(entry, components.TransformData{Pos: math.NewVec2(0,0), Rot: 0})
	components.Sprite.SetValue(entry, components.SpriteData{Image: assets.PlayerAsset0})
}

func (m *TestLevelScene) OnLeaveScene() {

}

func NewTestLevelScene(numEnemies int) TestLevelScene {
	return TestLevelScene{
		NumEnemies: numEnemies,
		text: "Hi there",
		world: donburi.NewWorld(),
		animSystem: *systems.NewAnimationSystem(),
		renderSystem: *systems.NewRenderSystem(),
	}
}

