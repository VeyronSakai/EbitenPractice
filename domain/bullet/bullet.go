package bullet

type Bullet struct {
	Id      Id
	Pos     Pos
	MoveDir MoveDir
	Speed   Speed
}

//func NewBullet() (*Bullet, error){
//	id := NewBulletId()
//
//}

func (b *Bullet) UpdatePos() {
	b.Pos.Value.X += b.MoveDir.Value.X * b.Speed.Value
	b.Pos.Value.Y += b.MoveDir.Value.Y * b.Speed.Value
}

func (b *Bullet) UpdateMoveDir(dir MoveDir) {
	b.MoveDir = dir
}
