package infrastructure

import (
	"EbitenSample/domain"
	"log"
)

type PlayerFactoryImpl struct {
}

func NewPlayerFactoryImpl() domain.PlayerFactory {
	return &PlayerFactoryImpl{}
}

func (factory *PlayerFactoryImpl) Create() (*domain.Player, error) {
	player, err := domain.NewPlayer()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return player, nil
}
