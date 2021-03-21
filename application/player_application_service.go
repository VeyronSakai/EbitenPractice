package application

import (
	"EbitenSample/domain/player"
	"log"
)

type PlayerApplicationService struct {
	playerRepository player.Repository
	playerFactory    player.Factory
}

func NewPlayerApplicationService(repository player.Repository, factory player.Factory) *PlayerApplicationService {
	return &PlayerApplicationService{repository, factory}
}

func (service *PlayerApplicationService) SpawnPlayer() (player.Id, error) {
	factory := service.playerFactory
	player, err := factory.Create()

	if err != nil {
		log.Fatal(err)
		return player.Id, err
	}

	repository := service.playerRepository
	repository.Save(player)

	return player.Id, nil
}

func (service *PlayerApplicationService) MovePlayer(id player.Id, dir player.MoveDir) {
	player, err := service.playerRepository.Find(id)

	if err != nil {
		log.Fatal(err)
	}

	player.UpdateMoveDir(dir)
	player.UpdatePos()
}

func (service *PlayerApplicationService) GetPlayerData(id player.Id) (*player.Player, error) {
	player, err := service.playerRepository.Find(id)

	if err != nil {
		log.Fatal(err)
	}

	return player, nil
}
