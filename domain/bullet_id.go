package domain

import "github.com/google/uuid"

type BulletId struct {
	Value string
}

func NewBulletId() BulletId {
	bulletId := BulletId{uuid.New().String()}
	return bulletId
}
