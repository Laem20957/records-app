package repository

import (
	"context"

	"github.com/Laem20957/records-app/internal/domain"
	psql "github.com/Laem20957/records-app/internal/repository/postgresql"
	"github.com/jmoiron/sqlx"
)

type IRepositoryAuthorizationMethods interface {
	CreateUserDB(ctx context.Context, user domain.Users) (int, error)
	GetUserDB(ctx context.Context, username, password string) (domain.Users, error)
	CreateTokenDB(ctx context.Context, token domain.RefreshSession) error
	GetTokenDB(ctx context.Context, token string) (domain.RefreshSession, error)
}

type IRepositoryRecordMethods interface {
	CreateRecordsDB(ctx context.Context, userId int, record domain.Records) (int, error)
	GetByIDRecordsDB(ctx context.Context, userId, id int) (domain.Records, error)
	GetAllRecordsDB(ctx context.Context) ([]domain.Records, error)
	DeleteRecordsDB(ctx context.Context, userId, id int) error
	UpdateRecordsDB(ctx context.Context, userId, id int, newNote domain.UpdateRecord) error
}

type RepositoryMethods struct {
	IRepositoryAuthorizationMethods
	IRepositoryRecordMethods
}

func RepositoryGetMethods(db *sqlx.DB) *RepositoryMethods {
	return &RepositoryMethods{
		IRepositoryAuthorizationMethods: psql.RepositoryGetAuth(db),
		IRepositoryRecordMethods:        psql.RepositoryGetRecord(db),
	}
}
