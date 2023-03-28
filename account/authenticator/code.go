package authenticator

import (
	"errors"
	"strings"
)

var (
	ErrEmptyAuthCode = errors.New("empty auth code supplied")
)

type AuthCode string

func NewAuthCode(c string) (AuthCode, error) {
	c = strings.TrimSpace(c)
	if c == "" {
		return "", ErrEmptyAuthCode
	}
	return AuthCode(c), nil
}

func (a AuthCode) String() string {
	return string(a)
}
