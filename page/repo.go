package page

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Add(ctx context.Context, p Page) error
	GetAll(ctx context.Context, id uuid.UUID) (*Page, error)
	GetAllForSite(ctx context.Context, siteID uuid.UUID) ([]Page, error)
}
