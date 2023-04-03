package site

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/internal"
	"github.com/samverrall/sitesmiths-api/site"
)

type CreatePayload struct {
	Name        string
	URL         string
	OwnerID     string
	Description string
}

func (s *Service) Create(ctx context.Context, p CreatePayload) error {
	siteName := site.NewName(p.Name)

	siteURL := site.NewURL(p.URL)

	ownerID, err := uuid.Parse(p.OwnerID)
	if err != nil {
		return internal.WrapErr(internal.ErrBadRequest, err)
	}

	description, err := site.NewDescription(p.Description)
	if err != nil {
		return errors.Join(internal.ErrBadRequest, err)
	}

	site := site.New(siteURL, siteName, description, ownerID)
	if err := s.repo.Add(ctx, site); err != nil {
		return internal.WrapErr(internal.ErrInternal, err)
	}

	return nil
}
