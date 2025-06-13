package scenes

import (
	c "github.com/BrianAnakPintar/ducktape/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TestLevelScene struct {
	NumEnemies int
	text string
}

func (m *TestLevelScene) GetName() string {
	return "TestLevel"
}

func (m *TestLevelScene) Update() {
	
}

func (m *TestLevelScene) Render(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, m.text)
}
func (m *TestLevelScene) HandleInput() {
	if ebiten.IsKeyPressed(c.SkipCutsceneKey) {	
		m.text = "Brian is goated"
	}
}

func NewTestLevelScene(numEnemies int) TestLevelScene {
	return TestLevelScene{
		NumEnemies: numEnemies,
		text: "Hi there",
	}
}

