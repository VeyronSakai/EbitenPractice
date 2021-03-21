package player

type Repository interface {
	Find(id Id) (*Player, error)
	Save(player *Player)
}
