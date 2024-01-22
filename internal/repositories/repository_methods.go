package repository

import (
	"context"

	domain "github.com/Laem20957/records-app/internal/domains"
	psql "github.com/Laem20957/records-app/internal/repositories/postgresql"
	"github.com/jmoiron/sqlx"
)

type RepositoryMethods struct {
	RepositoryAuthorizationMethods
	RepositoryRecordMethods
}

func RepositoryGetMethods(db *sqlx.DB) *RepositoryMethods {
	return &RepositoryMethods{
		RepositoryAuthorizationMethods: psql.RepositoryGetAuth(db),
		RepositoryRecordMethods:        psql.RepositoryGetRecord(db),
	}
}

type RepositoryAuthorizationMethods interface {
	CreateUsers(ctx context.Context, user domain.Users) (int, error)
	GetUsers(ctx context.Context, username, password string) (domain.Users, error)
	GetToken(ctx context.Context, token string) (domain.RefreshSession, error)
	CreateToken(ctx context.Context, token domain.RefreshSession) error
}

type RepositoryRecordMethods interface {
	CreateRecords(ctx context.Context, userId int, record domain.Record) (int, error)
	GetByIDRecords(ctx context.Context, userId, id int) (domain.Record, error)
	GetAllRecords(ctx context.Context, userId int) ([]domain.Record, error)
	DeleteRecords(ctx context.Context, userId, id int) error
	UpdateRecords(ctx context.Context, userId, id int, newNote domain.UpdateRecord) error
}
