package service

import (
	"context"

	"records-app/internal/domain"
	"records-app/internal/repository"
	settings "records-app/settings"

	"github.com/bluele/gcache"
)

type IServiceAuthorizationMethods interface {
	CreateUser(ctx context.Context, user domain.Users) (int, error)
	SignIn(ctx context.Context, input domain.Users) (string, string, error)
	GenerateToken(ctx context.Context, userId int) (string, string, error)
	TokenIsSigned(token string) (int, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type IServiceRecordMethods interface {
	GetAllRecords(ctx context.Context) ([]domain.Records, error)
	GetByIDRecords(ctx context.Context, recordId int) (domain.Records, error)
	CreateRecords(ctx context.Context) (int, error)
	UpdateRecords(ctx context.Context, newId int, newTitle string, newDescription string) (domain.Records, error)
	DeleteRecords(ctx context.Context, recordId int) (int, error)
}

type ServiceMethods struct {
	IServiceAuthorizationMethods
	IServiceRecordMethods
}

func ServiceGetMethods(settings *settings.Settings, cache gcache.Cache, repo *repository.RepositoryMethods) *ServiceMethods {
	return &ServiceMethods{
		IServiceAuthorizationMethods: nil, // ServiceGetAuth(settings, repo.IRepositoryAuthorizationMethods.(repository.RepositoryMethods)),
		IServiceRecordMethods:        ServiceGetRecords(settings, cache, repo.IRepositoryRecordMethods.(repository.RepositoryMethods)),
	}
}
