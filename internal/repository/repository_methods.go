package repository

import (
	"context"

	"records-app/internal/domain"

	"github.com/jinzhu/gorm"
)

type IRepositoryAuthorizationMethods interface {
	GetUserDB(ctx context.Context, userId int) (domain.Users, error)
	GetTokenDB(ctx context.Context, tokenId int) (domain.Tokens, error)
	CreateUserDB(ctx context.Context) (int, error)
	CreateTokenDB(ctx context.Context) (int, error)
}

type IRepositoryRecordMethods interface {
	GetAllRecordsDB(ctx context.Context) ([]domain.Records, error)
	GetByIDRecordsDB(ctx context.Context, recordId int) (domain.Records, error)
	CreateRecordsDB(ctx context.Context) (int, error)
	UpdateRecordsDB(ctx context.Context, newId int, newTitle string, newDescription string) (domain.Records, error)
	DeleteRecordsDB(ctx context.Context, recordId int) (int, error)
}

type RepositoryMethods struct {
	IRepositoryAuthorizationMethods
	IRepositoryRecordMethods
}

func RepositoryGetMethods(db *gorm.DB) *RepositoryMethods {
	return &RepositoryMethods{
		IRepositoryAuthorizationMethods: RepositoryGetAuth(db),
		IRepositoryRecordMethods:        RepositoryGetRecord(db),
	}
}
