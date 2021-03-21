package player

type Player struct {
	Id      Id
	Pos     Pos
	MoveDir MoveDir
	Speed   Speed
}

const initialSpeed float64 = 3.0

func NewPlayer() (*Player, error) {
	playerId := NewPlayerId()
	playerPos := Pos{}
	moveDir := MoveDir{}
	speed := Speed{Value: initialSpeed}
	player := &Player{playerId, playerPos, moveDir, speed}

	return player, nil
}

func (p *Player) UpdatePos() {
	p.Pos.Value.X += p.MoveDir.Value.X * p.Speed.Value
	p.Pos.Value.Y += p.MoveDir.Value.Y * p.Speed.Value
}

func (p *Player) UpdateMoveDir(moveDir MoveDir) {
	p.MoveDir = moveDir
}
