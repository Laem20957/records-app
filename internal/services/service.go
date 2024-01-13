package service

import (
	"context"

	config "github.com/Laem20957/records-app/configs"
	domain "github.com/Laem20957/records-app/internal/domains"
	repository "github.com/Laem20957/records-app/internal/repositories"
	"github.com/bluele/gcache"
)

// go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	SignIn(ctx context.Context, input domain.SignInInput) (string, string, error)
	GenerateTokens(ctx context.Context, userId int) (string, string, error)
	ParseToken(token string) (int, error)
	RefreshTokens(ctx context.Context, refreshToken string) (string, string, error)
}

type Note interface {
	Create(ctx context.Context, userId int, note domain.Note) (int, error)
	GetByID(ctx context.Context, userId, id int) (domain.Note, error)
	GetAll(ctx context.Context, userId int) ([]domain.Note, error)
	Delete(ctx context.Context, userId, id int) error
	Update(ctx context.Context, userId, id int, newNote domain.UpdateNote) error
}

type Service struct {
	Authorization
	Note
}

func GetService(config *config.Config, cache gcache.Cache, repos *repository.Repository) *Service {
	return &Service{
		Authorization: GetAuthService(config, repos.Authorization),
		Note:          GetNoteService(config, cache, repos.Note),
	}
}
