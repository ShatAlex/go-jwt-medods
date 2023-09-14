package repository

import (
	"context"
	"fmt"
	"time"

	tokens "github.com/ShatALex/TestTaskBackDev"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db *mongo.Collection
}

func NewAuthPostgres(db *mongo.Collection) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, user tokens.User) (string, error) {

	_, err := r.db.InsertOne(ctx, user)

	if err != nil {
		return "", fmt.Errorf("failed to create user, error: %v", err)
	}

	return user.Guid, nil
}

func (r *AuthPostgres) SetRefreshToken(ctx context.Context, refreshToken, guid string) error {

	filter := bson.D{{"guid", guid}}
	var user tokens.User
	err := r.db.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}

	user.RefreshToken = refreshToken
	user.ExpiresAt = time.Now().Add(720 * time.Hour)

	result, err := r.db.ReplaceOne(ctx, filter, user)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return err
	}

	return nil
}

func (r *AuthPostgres) ValidateRefreshToken(ctx context.Context, validateRefreshToken string, guid string) bool {

	filter := bson.D{{"guid", guid}}
	var user tokens.User
	err := r.db.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.RefreshToken), []byte(validateRefreshToken)); err == nil {
		return true
	}

	return false

}
