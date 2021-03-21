package domain

type PlayerFactory interface {
	Create() (*Player, error)
}
