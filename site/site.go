package site

import (
	"time"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/pkg/aggregate"
)

type Site struct {
	aggregate.Root

	ID          uuid.UUID
	URL         URL
	Name        Name
	Active      bool
	Description Description
	OwnerID     uuid.UUID
	Status      Status
	CreatedAt   time.Time
}

func New(url URL, name Name, desc Description, ownerID uuid.UUID) Site {
	return Site{
		ID:          uuid.New(),
		URL:         url,
		Name:        name,
		Description: desc,
		OwnerID:     ownerID,
		Active:      true,
		Status:      StatusDevelopment,
		CreatedAt:   time.Now().UTC(),
	}
}

func (s Site) IsStatusDevelopment() bool {
	return s.Status == StatusDevelopment
}
