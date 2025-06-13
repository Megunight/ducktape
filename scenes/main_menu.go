package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MainMenuScene struct {
}

func (m *MainMenuScene) GetName() string {
	return "MainMenu"
}

func (m *MainMenuScene) Update() {

}

func (m *MainMenuScene) Render(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World")
}
func (m *MainMenuScene) HandleInput() {

}

func NewMainMenu() MainMenuScene {
	return MainMenuScene{}
}
