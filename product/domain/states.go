package domain

type State int

const (
	ProductCreated State = iota
	ProductCreateError
)