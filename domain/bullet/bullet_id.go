package bullet

import "github.com/google/uuid"

type Id struct {
	Value string
}

func NewBulletId() Id {
	bulletId := Id{uuid.New().String()}
	return bulletId
}
