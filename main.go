package main

import (
	"EbitenSample/presentation"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &presentation.Game{}
	game.Initialize()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
