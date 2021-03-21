package domain

type Bullet struct {
	Id      BulletId
	Pos     BulletPos
	MoveDir BulletMoveDir
	Speed   BulletSpeed
}

//func NewBullet() (*Bullet, error){
//	id := NewBulletId()
//
//}

func (b *Bullet) UpdatePos() {
	b.Pos.Value.X += b.MoveDir.Value.X * b.Speed.Value
	b.Pos.Value.Y += b.MoveDir.Value.Y * b.Speed.Value
}

func (b *Bullet) UpdateMoveDir(dir BulletMoveDir) {
	b.MoveDir = dir
}
