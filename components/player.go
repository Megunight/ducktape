package components

import "github.com/yohamta/donburi"

type Player struct {
	Health  int
	Name 	string
}

var Players = donburi.NewComponentType[Player]()
