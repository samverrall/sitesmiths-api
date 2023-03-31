package page

import (
	"time"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/pkg/aggregate"
)

type Page struct {
	aggregate.Root

	ID        uuid.UUID
	SiteID    uuid.UUID
	Type      Type
	Heading   Heading
	Content   Content
	CreatedAt time.Time
}

func NewPage(id uuid.UUID, heading Heading, content Content, pageType Type, siteID uuid.UUID) Page {
	return Page{
		ID:        id,
		Type:      pageType,
		Heading:   heading,
		Content:   content,
		SiteID:    siteID,
		CreatedAt: time.Now().UTC(),
	}
}
