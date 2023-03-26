package repo

import (
	"context"

	"github.com/samverrall/sitesmiths-api/domain"
)

type Site interface {
	Add(ctx context.Context, s domain.Site) error
}
