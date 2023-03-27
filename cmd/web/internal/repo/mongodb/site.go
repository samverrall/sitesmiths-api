package mongodb

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/cmd/web/internal/repo/mongodb/internal"
	"github.com/samverrall/sitesmiths-api/site"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ site.Repo = &SiteRepo{}
)

type SiteRepo struct {
	collection *mongo.Collection
}

func NewSiteRepo(siteCollection *mongo.Collection) *SiteRepo {
	return &SiteRepo{
		collection: siteCollection,
	}
}

func (r *SiteRepo) Add(ctx context.Context, s site.Site) error {
	site := internal.Site{
		ID:        s.ID.String(),
		URL:       string(s.URL.String()),
		Name:      s.Name.String(),
		Active:    s.Active,
		OwnerID:   s.OwnerID.String(),
		Status:    s.Status.String(),
		CreatedAt: s.CreatedAt.UTC(),
	}
	_, err := r.collection.InsertOne(ctx, site)
	if err != nil {
		return err
	}

	return nil
}

func (r *SiteRepo) Get(ctx context.Context, id uuid.UUID) (*site.Site, error) {
	var s internal.Site
	filter := bson.M{"id": id.String()}
	err := r.collection.FindOne(ctx, filter).Decode(&s)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, site.ErrNotFound
		}
		return nil, err
	}

	return fromSiteModel(&s), nil
}

func fromSiteModel(model *internal.Site) *site.Site {
	return &site.Site{
		ID:        uuid.MustParse(model.ID),
		URL:       site.URL(model.URL),
		Name:      site.Name(model.Name),
		Active:    model.Active,
		OwnerID:   uuid.MustParse(model.OwnerID),
		Status:    site.Status(model.Status),
		CreatedAt: model.CreatedAt,
	}
}
