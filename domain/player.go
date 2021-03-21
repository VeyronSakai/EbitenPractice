package domain

type Player struct {
	Id      PlayerId
	Pos     PlayerPos
	MoveDir PlayerMoveDir
	Speed   PlayerSpeed
}

const initialSpeed float64 = 3.0

func NewPlayer() (*Player, error) {
	playerId := NewPlayerId()
	playerPos := PlayerPos{}
	moveDir := PlayerMoveDir{}
	speed := PlayerSpeed{Value: initialSpeed}
	player := &Player{playerId, playerPos, moveDir, speed}

	return player, nil
}

func (p *Player) UpdatePos() {
	p.Pos.Value.X += p.MoveDir.Value.X * p.Speed.Value
	p.Pos.Value.Y += p.MoveDir.Value.Y * p.Speed.Value
}

func (p *Player) UpdateMoveDir(moveDir PlayerMoveDir) {
	p.MoveDir = moveDir
}
