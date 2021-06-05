package main

import (
	"EbitenSample/infrastructure"
	"EbitenSample/presentation"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &presentation.Game{}

	playerRepository := infrastructure.NewPlayerRepositoryImpl()
	playerFactory := infrastructure.NewPlayerFactoryImpl()

	game.Initialize(playerRepository, playerFactory)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
