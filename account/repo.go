package account

import "context"

type Repo interface {
	Add(ctx context.Context, a Account) error
	GetByEmail(ctx context.Context, email Email) (Account, error)
}
