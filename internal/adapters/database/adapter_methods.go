package database

import (
	"context"

	"records-app/internal/adapters/database/schemas"

	"github.com/jinzhu/gorm"
)

type IAdapterAuthorizationMethods interface {
	GetUserDB(ctx context.Context, userId int) (schemas.Users, error)
	GetTokenDB(ctx context.Context, tokenId int) (schemas.Tokens, error)
	CreateUserDB(ctx context.Context) (int, error)
	CreateTokenDB(ctx context.Context) (int, error)
}

type IAdapterRecordMethods interface {
	GetAllRecordsDB(ctx context.Context) ([]schemas.Records, error)
	GetByIDRecordsDB(ctx context.Context, recordId int) (schemas.Records, error)
	CreateRecordsDB(ctx context.Context) (int, error)
	UpdateRecordsDB(ctx context.Context, newId int, newTitle, newDescription string) (schemas.Records, error)
	DeleteRecordsDB(ctx context.Context, recordId int) (int, error)
}

type AdapterMethods struct {
	IAdapterAuthorizationMethods
	IAdapterRecordMethods
}

func AdapterGetMethods(db *gorm.DB) *AdapterMethods {
	return &AdapterMethods{
		IAdapterAuthorizationMethods: NewGetAdapterAuth(db),
		IAdapterRecordMethods:        NewGetAdapterRecord(db),
	}
}
