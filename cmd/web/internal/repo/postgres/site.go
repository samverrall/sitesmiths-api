package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/site"
)

var (
	_ site.Repo = &SiteRepo{}
)

type SiteRepo struct {
	db *sql.DB
}

func NewSiteRepo(db *sql.DB) *SiteRepo {
	return &SiteRepo{
		db: db,
	}
}

func (s *SiteRepo) Add(ctx context.Context, site site.Site) error {
	return nil
}

func (s *SiteRepo) Get(ctx context.Context, id uuid.UUID) (*site.Site, error) {
	return nil, nil
}
