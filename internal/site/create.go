package site

import (
	"context"

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
		return internal.WrapErr(internal.ErrBadRequest, err)
	}

	site := site.New(siteURL, siteName, ownerID)
	if err := s.repo.Add(ctx, site); err != nil {
		return internal.WrapErr(internal.ErrInternal, err)
	}

	return nil
}
