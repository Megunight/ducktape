package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
    GetName() string
    Update()
    Render(screen *ebiten.Image)
    HandleInput()
}
