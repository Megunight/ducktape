package main

import (
	"log"

	"github.com/BrianAnakPintar/ducktape/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	currScene := scenes.GetSceneManager().GetCurrScene()
	currScene.HandleInput()
	currScene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scenes.GetSceneManager().GetCurrScene().Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func InitializeScenes() {
	mm := scenes.NewMainMenu()
	tl := scenes.NewTestLevelScene(1)

	availableScenes := [...]scenes.Scene{&mm, &tl}

	for _, scene := range availableScenes {
		scenes.GetSceneManager().RegisterScene(scene)
	}
}

func main() {
	InitializeScenes()
	scenes.GetSceneManager().SwitchSceneByName("TestLevel")

	ebiten.SetWindowSize(1000, 480)
	ebiten.SetWindowTitle("Ducktaped Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

