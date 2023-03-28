package internal

import (
	"time"
)

type Site struct {
	ID        string
	URL       string
	Name      string
	Active    bool
	OwnerID   string
	Status    string
	CreatedAt time.Time
}
