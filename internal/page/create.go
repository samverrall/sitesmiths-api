package page

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/internal"
	"github.com/samverrall/sitesmiths-api/page"
)

var (
	ErrInvalidSiteID    = errors.New("invalid site id supplied")
	ErrInvalidAccountID = errors.New("invalid account id supplied")
)

// CreatePayload defines a primative seriaisable payload with fields
// required for the Create page method.
type CreatePayload struct {
	SiteID    string
	AccountID string
	Content   string
	Heading   string
	Type      string
}

func (s *Service) Create(ctx context.Context, c CreatePayload) error {
	siteID, err := uuid.Parse(c.SiteID)
	if err != nil {
		return errors.Join(internal.ErrBadRequest, ErrInvalidSiteID)
	}

	accountID, err := uuid.Parse(c.AccountID)
	if err != nil {
		return errors.Join(internal.ErrBadRequest, ErrInvalidAccountID)
	}

	content, err := page.NewContent(c.Content)
	if err != nil {
		return errors.Join(internal.ErrBadRequest, err)
	}

	pageType, err := page.NewType(c.Type)
	if err != nil {
		return errors.Join(internal.ErrBadRequest, err)
	}

	page.New(uuid.New(), "", content, pageType, siteID, accountID)

	// TODO: check if site exists

	return nil
}
