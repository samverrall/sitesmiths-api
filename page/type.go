package page

import "errors"

type Type string

var pageTypes = map[string]struct{}{
	"about":   {},
	"home":    {},
	"contact": {},
}

func NewPageType(t string) (Type, error) {
	_, ok := pageTypes[t]
	if !ok {
		return "", errors.New("invalid page type")
	}
	return Type(t), nil
}
