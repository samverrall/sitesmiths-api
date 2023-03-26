package siteservice

import (
	"context"
	"fmt"

	"github.com/samverrall/sitesmiths-api/internal"
	"github.com/samverrall/sitesmiths-api/site"
)

type Service struct {
	repo site.Repo
}

func New(siteRepo site.Repo) *Service {
	return &Service{
		repo: siteRepo,
	}
}

func (s *Service) Create(ctx context.Context, site site.Site) error {
	if err := s.repo.Add(ctx, site); err != nil {
		return fmt.Errorf("%w: %s", internal.ErrInternal, err.Error())
	}
	return nil
}
