package page

import (
	"errors"

	"github.com/google/uuid"
)

type Page struct {
	ID        uuid.UUID
	SiteID    uuid.UUID
	Type      PageType
	Heading   PageHeading
	CreatedAt uuid.UUID
}

func NewPage(id uuid.UUID, t PageType, siteID uuid.UUID) Page {
	return Page{}
}

type PageType string

var pageTypes = map[string]struct{}{
	"about":   {},
	"home":    {},
	"contact": {},
}

func NewPageType(t string) (PageType, error) {
	_, ok := pageTypes[t]
	if !ok {
		return "", errors.New("invalid page type")
	}
	return PageType(t), nil

}

type PageHeading string

func NewPageHeading(h string) (PageHeading, error) {
	if h == "" {
		return "", errors.New("invalid page heading")
	}
	return PageHeading(h), nil

}
