package events

type Observer interface {
	Update(string)
	GetID() string
}
