package player

import "github.com/google/uuid"

type Id struct {
	Value string
}

func NewPlayerId() Id {
	playerId := Id{uuid.New().String()}
	return playerId
}
