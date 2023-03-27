package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	Provider  Provider
	Name      string // TODO: add object value types
	Email     string // TODO: add object value types
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(id uuid.UUID, provider Provider) Account {
	return Account{
		ID:        id,
		Provider:  provider,
		Active:    true,
		CreatedAt: time.Now(),
	}
}
