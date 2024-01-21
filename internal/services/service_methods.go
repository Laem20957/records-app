package services

import (
	"context"

	config "github.com/Laem20957/records-app/configs"
	domain "github.com/Laem20957/records-app/internal/domains"
	repository "github.com/Laem20957/records-app/internal/repositories"
	"github.com/bluele/gcache"
)

// go:generate mockgen -source=service.go -destination=mocks/mock.go

type ServiceMethods struct {
	ServiceAuthorizationMethods
	ServiceRecordMethods
}

func ServiceGetMethods(cfg *config.Config, cache gcache.Cache, repo *repository.RepositoryMethods) *ServiceMethods {
	return &ServiceMethods{
		ServiceAuthorizationMethods: ServiceGetAuth(cfg, repo.RepositoryAuthorizationMethods),
		ServiceRecordMethods:        ServiceGetRecords(cfg, cache, repo.RepositoryRecordMethods.(repository.RepositoryMethods)),
	}
}

type ServiceAuthorizationMethods interface {
	CreateUsers(ctx context.Context, user domain.Users) (int, error)
	SignIn(ctx context.Context, input domain.SignInInput) (string, string, error)
	GenerateTokens(ctx context.Context, userId int) (string, string, error)
	ParseToken(token string) (int, error)
	RefreshTokens(ctx context.Context, refreshToken string) (string, string, error)
}

type ServiceRecordMethods interface {
	CreateRecords(ctx context.Context, userId int, note domain.Record) (int, error)
	GetByIDRecords(ctx context.Context, userId, id int) (domain.Record, error)
	GetAllRecords(ctx context.Context, userId int) ([]domain.Record, error)
	DeleteRecords(ctx context.Context, userId, id int) error
	UpdateRecords(ctx context.Context, userId, id int, record domain.UpdateRecord) error
}
