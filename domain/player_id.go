package domain

import "github.com/google/uuid"

type PlayerId struct {
	Value string
}

func NewPlayerId() PlayerId {
	playerId := PlayerId{uuid.New().String()}
	return playerId
}
