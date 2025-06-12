package main

import (
	"log"

	"github.com/BrianAnakPintar/ducktape/SceneManager"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	currScene := scenemanager.GetInstance().GetCurrScene()
	currScene.HandleInput()
	currScene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scenemanager.GetInstance().GetCurrScene().Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func InitializeScenes() {
	mm := scenemanager.NewMainMenu()
	scenemanager.GetInstance().RegisterScene(&mm)
}

func main() {
	InitializeScenes()
	scenemanager.GetInstance().SwitchSceneByName("MainMenu")

	ebiten.SetWindowSize(1000, 480)
	ebiten.SetWindowTitle("Circuit Game!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
