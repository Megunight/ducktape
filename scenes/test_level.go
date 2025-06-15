package scenes

import (
	"github.com/BrianAnakPintar/ducktape/archetypes"
	"github.com/BrianAnakPintar/ducktape/components"
	c "github.com/BrianAnakPintar/ducktape/constants"
	"github.com/BrianAnakPintar/ducktape/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
)

type TestLevelScene struct {
	NumEnemies int
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
	t.renderSystem.Draw(t.world, screen)
}

func (t *TestLevelScene) HandleInput() {
	if ebiten.IsKeyPressed(c.MoveLeftKey) {	
		query := donburi.NewQuery(filter.Contains(components.Player, components.Transform, components.Sprite))
		if entry, ok := query.First(t.world); ok {
			components.Transform.Get(entry).Pos.X -= 1
		}
	} else if ebiten.IsKeyPressed(c.MoveRightKey) {
		query := donburi.NewQuery(filter.Contains(components.Player, components.Transform, components.Sprite))
		if entry, ok := query.First(t.world); ok {
			components.Transform.Get(entry).Pos.X += 1
		}
	}
}

func (t *TestLevelScene) OnEnterScene() {
	archetypes.NewPlayer(t.world, math.NewVec2(0,10))
}

func (m *TestLevelScene) OnLeaveScene() {

}

func NewTestLevelScene(numEnemies int) TestLevelScene {
	return TestLevelScene{
		NumEnemies: numEnemies,
		world: donburi.NewWorld(),
		animSystem: *systems.NewAnimationSystem(),
		renderSystem: *systems.NewRenderSystem(),
	}
}

