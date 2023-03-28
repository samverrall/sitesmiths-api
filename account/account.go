package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	Provider  Provider
	Name      Name
	Email     Email
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(id uuid.UUID, name Name, email Email, provider Provider) Account {
	return Account{
		ID:        id,
		Name:      name,
		Email:     email,
		Provider:  provider,
		Active:    true,
		CreatedAt: time.Now(),
	}
}
