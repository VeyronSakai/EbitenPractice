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
	p, err := factory.Create()

	if err != nil {
		log.Fatal(err)
		return p.Id, err
	}

	repository := service.playerRepository
	repository.Save(p)

	return p.Id, nil
}

func (service *PlayerApplicationService) MovePlayer(id player.Id, dir player.MoveDir) {
	p, err := service.playerRepository.Find(id)

	if err != nil {
		log.Fatal(err)
	}

	p.UpdateMoveDir(dir)
	p.UpdatePos()
}

func (service *PlayerApplicationService) GetPlayerData(id player.Id) (*player.Player, error) {
	p, err := service.playerRepository.Find(id)

	if err != nil {
		log.Fatal(err)
	}

	return p, nil
}
