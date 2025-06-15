package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	PlayerAsset0 *ebiten.Image = LoadAsset("assets/characters/player/player_idle00.png")
	PlayerAsset1 = LoadAsset("assets/characters/player/player_idle01.png")
	PlayerAsset2 = LoadAsset("assets/characters/player/player_idle02.png")
	PlayerAsset3 = LoadAsset("assets/characters/player/player_idle03.png")
	PlayerAsset4 = LoadAsset("assets/characters/player/player_idle04.png")
	PlayerAsset5 = LoadAsset("assets/characters/player/player_idle05.png")
)

func LoadAsset(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		log.Fatalf("Error loading asset %s: %v", path, err)
		return nil
	}

	return img
}
