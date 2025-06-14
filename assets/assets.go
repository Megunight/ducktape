package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	PlayerAsset = LoadAsset("")
)

func LoadAsset(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		log.Fatalf("Error loading asset %s: %v", path, err)
		return nil
	}

	return img
}
