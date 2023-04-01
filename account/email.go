package account

import (
	"errors"
	"strings"
)

var (
	ErrEmptyEmail = errors.New("empty email supplied")
)

type Email string

func NewEmail(e string) (Email, error) {
	e = strings.TrimSpace(e)
	if e == "" {
		return "", ErrEmptyEmail
	}
	return Email(e), nil
}

func (e Email) String() string {
	return string(e)
}
