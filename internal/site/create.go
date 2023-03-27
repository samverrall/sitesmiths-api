package siteservice

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/internal"
	"github.com/samverrall/sitesmiths-api/site"
)

type CreatePayload struct {
	Name    string
	URL     string
	OwnerID string
}

func (s *Service) Create(ctx context.Context, p CreatePayload) error {
	siteName := site.NewName(p.Name)

	siteURL := site.NewURL(p.URL)

	ownerID, err := uuid.Parse(p.OwnerID)
	if err != nil {
		return fmt.Errorf("%w: %s", internal.ErrBadRequest, err.Error())
	}

	site := site.New(siteURL, siteName, ownerID)
	if err := s.repo.Add(ctx, site); err != nil {
		return fmt.Errorf("%w: %s", internal.ErrInternal, err.Error())
	}

	return nil
}
