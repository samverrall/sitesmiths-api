package site

import (
	"time"

	"github.com/google/uuid"
)

type Site struct {
	ID        uuid.UUID
	URL       SiteURL
	Name      SiteName
	Active    bool
	OwnerID   uuid.UUID
	CreatedAt time.Time
}

func NewSite(url SiteURL, name SiteName, ownerID uuid.UUID) Site {
	return Site{
		URL:     url,
		Name:    name,
		OwnerID: ownerID,
	}
}
