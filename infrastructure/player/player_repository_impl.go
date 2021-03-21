package player

import (
	"EbitenSample/domain/player"
	"errors"
)

type RepositoryImpl struct {
	db map[player.Id]*player.Player
}

func NewPlayerRepositoryImpl() player.Repository {
	return &RepositoryImpl{db: make(map[player.Id]*player.Player)}
}

func (p RepositoryImpl) Find(id player.Id) (*player.Player, error) {
	if player, ok := p.db[id]; ok {
		return player, nil
	}

	return nil, errors.New("player doesn't exist")
}

func (p RepositoryImpl) Save(player *player.Player) {
	p.db[player.Id] = player
}
