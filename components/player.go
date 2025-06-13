package components

import "github.com/yohamta/donburi"

type PlayerData struct {
	Health  int
	Name 	string
}

var Player = donburi.NewComponentType[PlayerData]()
