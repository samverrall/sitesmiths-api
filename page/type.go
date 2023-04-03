package page

import (
	"errors"
)

var (
	ErrInvalidPageType = errors.New("invalid page type supplied")
)

type Type string

const (
	TypeHome    Type = "home"
	TypeAbout   Type = "about"
	TypeContact Type = "contact"
)

var pageTypes = map[string]struct{}{
	TypeHome.String():    {},
	TypeAbout.String():   {},
	TypeContact.String(): {},
}

func NewType(t string) (Type, error) {
	_, ok := pageTypes[t]
	if !ok {
		return "", ErrInvalidPageType
	}

	return Type(t), nil
}

func (t Type) String() string {
	return string(t)
}
