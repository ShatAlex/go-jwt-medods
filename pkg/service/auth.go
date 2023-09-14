package service

import (
	"context"
	"crypto/sha512"
	"errors"
	"fmt"
	"math/rand"
	"time"

	tokens "github.com/ShatALex/TestTaskBackDev"
	"github.com/ShatALex/TestTaskBackDev/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt           = "fW52sz01fAPLGgZ"
	signingKey     = "asdwSBd#aLtN#ad14"
	accessTokenTTL = 2 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	Guid string `json:"guid"`
}

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (s *AuthService) CreateUser(ctx context.Context, user tokens.SignUpUser) (string, error) {

	uuid := uuid.New().String()

	userInsert := tokens.User{
		Username: user.Username,
		Password: GenerateSHA512Hash(user.Password),
		Guid:     uuid,
	}

	return s.rep.CreateUser(ctx, userInsert)
}

func (s *AuthService) GenerateTokens(ctx context.Context, guid string) (string, string, error) {

	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
		},
		Guid: guid,
	})

	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", nil
	}

	refresDb, _ := hashRefreshToken(refreshToken)
	if err = s.rep.SetRefreshToken(ctx, refresDb, guid); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims aren't of type *tokenCLaims")
	}

	return claims.Guid, nil
}

func (s *AuthService) ValidateRefreshToken(ctx context.Context, refreshToken string, guid string) bool {
	return s.rep.ValidateRefreshToken(ctx, refreshToken, guid)
}

func GenerateSHA512Hash(guid string) string {
	hash := sha512.New()
	hash.Write([]byte(fmt.Sprint(guid)))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func newRefreshToken() (string, error) {
	token := make([]byte, 32)

	source := rand.NewSource(time.Now().Unix())
	rand := rand.New(source)

	if _, err := rand.Read(token); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", token), nil
}

func hashRefreshToken(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
