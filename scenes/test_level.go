package scenes

import (
	"image/color"

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
	physicsSystem systems.PhysicsSystem
	gravitySystem systems.GravitySystem

	playerQuery donburi.Query
}

func (t *TestLevelScene) GetName() string {
	return "TestLevel"
}

func (t *TestLevelScene) Update() {
	t.animSystem.Update(t.world)
	t.gravitySystem.Update(t.world)
	t.physicsSystem.Update(t.world)
}

func (t *TestLevelScene) Render(screen *ebiten.Image) {
	screen.Fill(color.White)
	t.renderSystem.Draw(t.world, screen)
}

func (t *TestLevelScene) HandleInput() {
	t.HandlePlayerMovement()
}

func (t *TestLevelScene) OnEnterScene() {
	archetypes.NewPlayer(t.world, math.NewVec2(0,100))
}

func (t *TestLevelScene) OnLeaveScene() {

}

func NewTestLevelScene(numEnemies int) TestLevelScene {
	return TestLevelScene{
		NumEnemies: numEnemies,
		world: donburi.NewWorld(),
		animSystem: *systems.NewAnimationSystem(),
		renderSystem: *systems.NewRenderSystem(),
		physicsSystem: *systems.NewPhysicsSystem(),
		gravitySystem: *systems.NewGravitySystem(),
		playerQuery: *donburi.NewQuery(filter.Contains(components.Player, components.Velocity, components.Transform)),
	}
}

func (t *TestLevelScene) HandlePlayerMovement() {
	entry, ok := t.playerQuery.First(t.world)
	if !ok {
		return
	}

	velocity := components.Velocity.Get(entry)

	const moveSpeed = 5.0
	velocity.PosVelocity.X = 0

	if ebiten.IsKeyPressed(c.MoveLeftKey) {
		velocity.PosVelocity.X -= moveSpeed
	}
	if ebiten.IsKeyPressed(c.MoveRightKey) {
		velocity.PosVelocity.X += moveSpeed
	}

	if ebiten.IsKeyPressed(c.JumpKey) {
		if components.Jump.Get(entry).JumpsLeft > 0 {
			const jumpVelocity = -c.JumpForce
			velocity.PosVelocity.Y = jumpVelocity
			components.Jump.Get(entry).JumpsLeft -= 1
		}
	}
}
