package service

import (
	"context"

	tokens "github.com/ShatALex/TestTaskBackDev"
	"github.com/ShatALex/TestTaskBackDev/pkg/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user tokens.SignUpUser) (string, error)
	GenerateTokens(ctx context.Context, guid string) (string, string, error)
	TakeGuidByRefToken(ctx context.Context, refreshToken string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
