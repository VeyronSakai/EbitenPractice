package main

import (
	"EbitenSample/infrastructure"
	"EbitenSample/presentation"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	playerRepository := infrastructure.NewPlayerRepositoryImpl()
	playerFactory := infrastructure.NewPlayerFactoryImpl()

	game, err := presentation.NewGame(playerRepository, playerFactory)
	if err != nil {
		log.Fatal(err)
	}

	presentation.InitializeView()
	if err = ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
