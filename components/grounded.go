package components

import "github.com/yohamta/donburi"

type JumpData struct {
	MaxJumps  int
	JumpsLeft int
}

var Jump = donburi.NewComponentType[JumpData]()
