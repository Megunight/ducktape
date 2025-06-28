package scenes

import (
	"image/color"
	"log"

	"github.com/BrianAnakPintar/ducktape/archetypes"
	"github.com/BrianAnakPintar/ducktape/components"
	c "github.com/BrianAnakPintar/ducktape/constants"
	"github.com/BrianAnakPintar/ducktape/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
)

type TestLevelScene struct {
	NumEnemies int
	world donburi.World
	levelMap *tiled.Map

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

func (t *TestLevelScene) HandleLayer(layer *tiled.Layer) {
	if layer.Tiles == nil {
		return
	}

	if layer.Name == c.PlayerSpawnLayer {
		for y := range t.levelMap.Height {
			for x := range t.levelMap.Width {
				tile := layer.Tiles[y*t.levelMap.Width + x]
				if tile == nil {
					continue
				}
				archetypes.NewPlayer(t.world, math.NewVec2(0,100))
			}
		}
	} else if layer.Name == c.CollisionLayer {
		for y := range t.levelMap.Height {
			for x := range t.levelMap.Width {
				tile := layer.Tiles[y*t.levelMap.Width + x]
				if tile == nil {
					continue
				}
			}
		}
	}	
}

func (t *TestLevelScene) LoadMap(path string) {
	levelMap, err := tiled.LoadFile(path)
	if err != nil {
		log.Fatalf("Error loading map: %v", err)
	}
	t.levelMap = levelMap

	for _, layer := range t.levelMap.Layers {
		t.HandleLayer(layer)
	}
}

func (t *TestLevelScene) OnEnterScene() {
	t.LoadMap(c.SpawnMapPath)
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
