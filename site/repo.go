package site

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("could not find site")
)

type Repo interface {
	Add(ctx context.Context, s Site) error
	Get(ctx context.Context, id uuid.UUID) (*Site, error)
}
