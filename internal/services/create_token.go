package services

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"time"

	config "github.com/Laem20957/records-app/configs"
	domain "github.com/Laem20957/records-app/internal/domains"
	repository "github.com/Laem20957/records-app/internal/repositories"
	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type ServiceAuth struct {
	config *config.Config
	repo   repository.IRepositoryAuthorizationMethods
}

func ServiceGetAuth(cfg *config.Config, repo repository.IRepositoryAuthorizationMethods) *ServiceAuth {
	return &ServiceAuth{cfg, repo}
}

func generatePasswordHash(s *config.Config, password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.Salt)))
}

func (s *ServiceAuth) GenerateToken(ctx context.Context, userId int) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.config.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	accessToken, err := token.SignedString([]byte(s.config.SigningKey))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.RefreshToken(ctx, accessToken)
	if err != nil {
		return "", "", err
	}

	if err := s.repo.CreateTokenDB(ctx, domain.RefreshSession{
		UserID:    userId,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.config.RefreshTokenTTL),
	}); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

func (s *ServiceAuth) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	refreshSession, err := s.repo.GetTokenDB(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	if refreshSession.ExpiresAt.Unix() < time.Now().Unix() {
		buffer := make([]byte, 32)
		srcRandomNumbers := rand.NewSource(time.Now().Unix())
		generatorRandomNumbers := rand.New(srcRandomNumbers)

		if _, err := generatorRandomNumbers.Read(buffer); err != nil {
			return "", err
		}

		newRefreshToken := fmt.Sprintf("%x", buffer)

		refreshSession.Token = newRefreshToken
		refreshSession.ExpiresAt = time.Now().Add(s.config.RefreshTokenTTL)
		if err := s.repo.CreateTokenDB(ctx, refreshSession); err != nil {
			return "", err
		}
		return newRefreshToken, err
	}
	return refreshToken, err
}

func (s *ServiceAuth) TokenIsSigned(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.config.SigningKey), nil
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

// func (s *ServiceAuth) RefreshToken() (string, error) {
// 	b := make([]byte, 32)
//
// 	ns := rand.NewSource(time.Now().Unix())
// 	r := rand.New(ns)
//
// 	if _, err := r.Read(b); err != nil {
// 		return "", err
// 	}
//
// 	return fmt.Sprintf("%x", b), nil
// }

// func (s *ServiceAuth) TokenIsExpired(ctx context.Context, refreshToken string) (string, string, error) {
// 	session, err := s.repo.GetTokenDB(ctx, refreshToken)
// 	if err != nil {
// 		return "", "", err
// 	}
//
// 	if session.ExpiresAt.Unix() < time.Now().Unix() {
// 		return "", "", errors.New("refresh token expired")
// 	}
//
// 	return s.GenerateToken(ctx, session.UserID)
// }

func (s *ServiceAuth) CreateUser(ctx context.Context, user domain.Users) (int, error) {
	user.Password = generatePasswordHash(s.config, user.Password)
	return s.repo.CreateUserDB(ctx, user)
}

func (s *ServiceAuth) SignIn(ctx context.Context, input domain.SignInInput) (string, string, error) {
	user, err := s.repo.GetUserDB(ctx, input.Username, generatePasswordHash(s.config, input.Password))
	if err != nil {
		return "", "", err
	}

	return s.GenerateToken(ctx, user.Id)
}
