package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/pkg/aggregate"
)

type Account struct {
	aggregate.Root

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
