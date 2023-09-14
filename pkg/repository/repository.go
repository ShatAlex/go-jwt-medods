package repository

import (
	"context"

	tokens "github.com/ShatALex/TestTaskBackDev"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(ctx context.Context, user tokens.User) (string, error)
	SetRefreshToken(ctx context.Context, refreshToken, guid string) error
	TakeGuidByRefToken(ctx context.Context, refreshToken string) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Database, collection string) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db.Collection(collection)),
	}
}
