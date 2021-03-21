package domain

type PlayerRepository interface {
	Find(id PlayerId) (*Player, error)
	Save(player *Player)
}