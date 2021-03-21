package domain

type PlayerFactory interface {
	Create(playerId PlayerId) (*Player, error)
}
