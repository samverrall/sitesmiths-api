package account

import "errors"

type Provider string

var (
	ErrInvalidProvider = errors.New("invalid provider supplied")
)

const (
	GoogleProvider Provider = "google"
)

func NewProvider(provider string) (Provider, error) {
	if provider != GoogleProvider.String() { // only support google as of now
		return "", ErrInvalidProvider
	}
	return Provider(provider), nil
}

func (p Provider) String() string {
	return string(p)
}
