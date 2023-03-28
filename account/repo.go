package account

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("no account found")
)

type Repo interface {
	Add(ctx context.Context, a Account) error
	GetByEmail(ctx context.Context, email Email) (Account, error)
}
