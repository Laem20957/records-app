package service

import (
	"context"

	config "github.com/Laem20957/records-app/configuration"
	"github.com/Laem20957/records-app/internal/domain"
	"github.com/Laem20957/records-app/internal/repository"
	"github.com/bluele/gcache"
)

type IServiceAuthorizationMethods interface {
	CreateUser(ctx context.Context, user domain.Users) (int, error)
	SignIn(ctx context.Context, input domain.SignInInput) (string, string, error)
	GenerateToken(ctx context.Context, userId int) (string, string, error)
	TokenIsSigned(token string) (int, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type IServiceRecordMethods interface {
	CreateRecords(ctx context.Context, userId int, note domain.Record) (int, error)
	GetByIDRecords(ctx context.Context, userId, id int) (domain.Record, error)
	GetAllRecords(ctx context.Context, userId int) ([]domain.Record, error)
	DeleteRecords(ctx context.Context, userId, id int) error
	UpdateRecords(ctx context.Context, userId, id int, record domain.UpdateRecord) error
}

type ServiceMethods struct {
	IServiceAuthorizationMethods
	IServiceRecordMethods
}

func ServiceGetMethods(cfg *config.Config, cache gcache.Cache, repo *repository.RepositoryMethods) *ServiceMethods {
	return &ServiceMethods{
		IServiceAuthorizationMethods: ServiceGetAuth(cfg, repo.IRepositoryAuthorizationMethods.(repository.RepositoryMethods)),
		IServiceRecordMethods:        ServiceGetRecords(cfg, cache, repo.IRepositoryRecordMethods.(repository.RepositoryMethods)),
	}
}
