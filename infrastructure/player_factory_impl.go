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

func (factory *PlayerFactoryImpl) Create(playerId domain.PlayerId) (*domain.Player, error) {
	player, err := domain.NewPlayer(playerId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return player, nil
}
