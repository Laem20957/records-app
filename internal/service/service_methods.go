package service

import (
	"context"

	"records-app/internal/adapters/database"
	"records-app/internal/adapters/database/schemas"
	"records-app/settings"

	"github.com/bluele/gcache"
)

type IServiceAuthorizationMethods interface {
	CreateUser(ctx context.Context, user schemas.Users) (int, error)
	SignIn(ctx context.Context, input schemas.Users) (string, string, error)
	GenerateToken(ctx context.Context, userId int) (string, string, error)
	TokenIsSigned(token string) (int, error)
	RefreshToken(ctx context.Context, tokenId int) (string, error)
}

type IServiceRecordMethods interface {
	GetAllRecords(ctx context.Context) ([]schemas.Records, error)
	GetByIDRecords(ctx context.Context, recordId int) (schemas.Records, error)
	CreateRecords(ctx context.Context) (int, error)
	UpdateRecords(ctx context.Context, newId int, newTitle string, newDescription string) (schemas.Records, error)
	DeleteRecords(ctx context.Context, recordId int) (int, error)
}

type ServiceMethods struct {
	IServiceAuthorizationMethods
	IServiceRecordMethods
}

func ServiceGetMethods(settings *settings.Settings, cache *gcache.Cache, db *database.AdapterMethods) *ServiceMethods {
	return &ServiceMethods{
		IServiceAuthorizationMethods: NewGetServiceAuth(settings, db.IAdapterAuthorizationMethods.(database.AdapterMethods)),
		IServiceRecordMethods:        NewGetServiceRecord(settings, *cache, db.IAdapterRecordMethods.(database.AdapterMethods)),
	}
}

//! Добавил указатель cache *gcache.Cache и *cache
