package application

import (
	"EbitenSample/domain"
	"log"
)

type PlayerApplicationService struct {
	playerRepository domain.PlayerRepository
	playerFactory    domain.PlayerFactory
}

func NewPlayerApplicationService(repository domain.PlayerRepository, factory domain.PlayerFactory) *PlayerApplicationService {
	return &PlayerApplicationService{repository, factory}
}

func (service *PlayerApplicationService) SpawnPlayer() (domain.PlayerId, error) {
	factory := service.playerFactory
	playerId := domain.NewPlayerId()
	player, err := factory.Create(playerId)

	if err != nil {
		log.Fatal(err)
		return domain.PlayerId{}, err
	}

	repository := service.playerRepository
	repository.Save(player)

	return playerId, nil
}

func (service *PlayerApplicationService) MovePlayer(id domain.PlayerId, dir domain.PlayerMoveDir){
	player, err := service.playerRepository.Find(id)

	if err != nil{
		log.Fatal(err)
	}

	player.UpdateMoveDir(dir)
	player.UpdatePos()
}

func (service *PlayerApplicationService) GetPlayerData(id domain.PlayerId) (*domain.Player, error){
	player, err := service.playerRepository.Find(id)

	if err != nil{
		log.Fatal(err)
	}

	return player, nil
}