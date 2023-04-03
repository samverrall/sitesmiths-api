package internal

import (
	"time"
)

type Site struct {
	ID          string
	URL         string
	Name        string
	Description string
	Active      bool
	OwnerID     string
	Status      string
	CreatedAt   time.Time
}
