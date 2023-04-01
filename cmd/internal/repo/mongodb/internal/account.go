package internal

import "time"

type Account struct {
	ID        string
	Provider  string
	Name      string
	Email     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
