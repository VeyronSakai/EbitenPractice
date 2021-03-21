package player

type Factory interface {
	Create() (*Player, error)
}
