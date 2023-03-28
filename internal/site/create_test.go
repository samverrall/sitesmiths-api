package site

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/internal"
	"github.com/samverrall/sitesmiths-api/site"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	ctx := context.Background()
	repo := site.NewMockRepo(ctrl)

	repo.EXPECT().Add(ctx, gomock.Any()).Return(nil)

	svc := NewService(repo)

	t.Run("successful", func(t *testing.T) {
		err := svc.Create(ctx, CreatePayload{
			Name:    "Site",
			URL:     "site.com",
			OwnerID: uuid.NewString(),
		})
		if err != nil {
			t.Errorf("want <nil> error got: %v", err)
		}
	})

	t.Run("invalid owner guid supplied", func(t *testing.T) {
		err := svc.Create(ctx, CreatePayload{
			OwnerID: "invalidGUID",
		})
		if !errors.Is(err, internal.ErrBadRequest) {
			t.Errorf("want bad request error, got %v", err)
		}
		if err == nil {
			t.Errorf("want err, got <nil>")
		}
	})

	t.Run("repo failure", func(t *testing.T) {
		repo.EXPECT().Add(ctx, gomock.Any()).Return(errors.New("repo error"))

		err := svc.Create(ctx, CreatePayload{
			Name:    "Site",
			URL:     "site.com",
			OwnerID: uuid.NewString(),
		})
		if !errors.Is(err, internal.ErrInternal) {
			t.Errorf("want internal error, got %v", err)
		}
		if err == nil {
			t.Errorf("want err, got <nil>")
		}
	})
}
