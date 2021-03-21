package infrastructure

import (
	"EbitenSample/domain"
	"errors"
)

type PlayerRepositoryImpl struct {
	db map[domain.PlayerId]*domain.Player
}

func NewPlayerRepositoryImpl() domain.PlayerRepository {
	return &PlayerRepositoryImpl{db: make(map[domain.PlayerId]*domain.Player)}
}

func (p PlayerRepositoryImpl) Find(id domain.PlayerId) (*domain.Player, error){
	if player, ok := p.db[id]; ok{
		return player, nil
	}

	return nil, errors.New("player doesn't exist")
}

func (p PlayerRepositoryImpl) Save(player *domain.Player) {
	p.db[player.Id] = player
}
