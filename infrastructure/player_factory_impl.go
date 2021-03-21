package infrastructure

import (
	"EbitenSample/domain/player"
	"log"
)

type FactoryImpl struct {
}

func NewPlayerFactoryImpl() player.Factory {
	return &FactoryImpl{}
}

func (factory *FactoryImpl) Create() (*player.Player, error) {
	player, err := player.NewPlayer()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return player, nil
}
