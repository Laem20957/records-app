package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"records-app/internal/adapters/database"
	"records-app/internal/adapters/database/schemas"
	"records-app/settings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type ServiceAuth struct {
	settings *settings.Settings
	db       database.IAdapterAuthorizationMethods
}

func NewGetServiceAuth(settings *settings.Settings, db database.IAdapterAuthorizationMethods) *ServiceAuth {
	return &ServiceAuth{settings, db}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("sorbmobe")))
}

func (s *ServiceAuth) GenerateToken(ctx context.Context, userId int) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.settings.TTLToken).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	accessToken, err := token.SignedString([]byte(s.settings.SignKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.RefreshToken(ctx, 1)
	if err != nil {
		return "", "", err
	}

	// 	return accessToken, refreshToken

	if _, err := s.db.CreateTokenDB(ctx); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}

func (s *ServiceAuth) RefreshToken(ctx context.Context, tokenId int) (string, error) {
	refreshSession, err := s.db.GetTokenDB(ctx, tokenId)
	if err != nil {
		return "", err
	}

	if refreshSession.Expiration.Unix() < time.Now().Unix() {
		buffer := make([]byte, 32)
		srcRandomNumbers := rand.NewSource(time.Now().Unix())
		generatorRandomNumbers := rand.New(srcRandomNumbers)

		if _, err := generatorRandomNumbers.Read(buffer); err != nil {
			return "", err
		}

		refreshToken := fmt.Sprintf("%x", buffer)

		refreshSession.Token = refreshToken
		refreshSession.Expiration = time.Now().Add(s.settings.TTLRefreshToken)

		if _, err := s.db.CreateTokenDB(ctx); err != nil {
			return "", err
		}
		return refreshToken, err
	}
	return refreshSession.Token, err
}

func (s *ServiceAuth) TokenIsSigned(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.settings.SignKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *ServiceAuth) CreateUser(ctx context.Context, user schemas.Users) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.db.CreateUserDB(ctx)
}

func (s *ServiceAuth) SignIn(ctx context.Context, user schemas.Users) (string, string, error) {
	user, err := s.db.GetUserDB(ctx, user.ID)
	if err != nil {
		return "", "", err
	}
	return s.GenerateToken(ctx, user.ID)
}
