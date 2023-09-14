package service

import (
	"context"

	tokens "github.com/ShatALex/TestTaskBackDev"
	"github.com/ShatALex/TestTaskBackDev/pkg/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user tokens.SignUpUser) (string, error)
	GenerateTokens(ctx context.Context, guid string) (string, string, error)
	ParseToken(accessToken string) (string, error)
	ValidateRefreshToken(ctx context.Context, refreshToken string, guid string) bool
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
