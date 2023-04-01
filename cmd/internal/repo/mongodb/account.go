package mongodb

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samverrall/sitesmiths-api/account"
	"github.com/samverrall/sitesmiths-api/cmd/internal/repo/mongodb/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const accountCollection = "accounts"

var (
	_ account.Repo = &AccountRepo{}
)

type AccountRepo struct {
	collection *mongo.Collection
}

func NewAccountRepo(db *mongo.Database) *AccountRepo {
	accountsCollection := db.Collection(accountCollection)

	return &AccountRepo{
		collection: accountsCollection,
	}
}

func (r *AccountRepo) Add(ctx context.Context, a account.Account) error {
	newAcc := internal.Account{
		ID:        a.ID.String(),
		Provider:  a.Provider.String(),
		Name:      a.Name.String(),
		Email:     a.Email.String(),
		Active:    a.Active,
		CreatedAt: a.CreatedAt.UTC(),
	}
	_, err := r.collection.InsertOne(ctx, newAcc)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepo) GetByEmail(ctx context.Context, email account.Email) (account.Account, error) {
	var a internal.Account
	filter := bson.M{"email": email.String()}
	err := r.collection.FindOne(ctx, filter).Decode(&a)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return account.Account{}, account.ErrNotFound
		}
		return account.Account{}, err
	}
	return newAccountFromModel(a), nil
}

func newAccountFromModel(model internal.Account) account.Account {
	return account.Account{
		ID:        uuid.MustParse(model.ID),
		Provider:  account.Provider(model.Provider),
		Name:      account.Name(model.Name),
		Email:     account.Email(model.Email),
		Active:    model.Active,
		CreatedAt: model.CreatedAt,
	}
}
