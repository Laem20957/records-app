package repository

import (
	"context"

	domain "github.com/Laem20957/records-app/internal/domains"
	"github.com/Laem20957/records-app/internal/repositories/psql"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	GetUser(ctx context.Context, username, password string) (domain.User, error)
	GetToken(ctx context.Context, token string) (domain.RefreshSession, error)
	CreateToken(ctx context.Context, token domain.RefreshSession) error
}

type Note interface {
	Create(ctx context.Context, userId int, note domain.Note) (int, error)
	GetByID(ctx context.Context, userId, id int) (domain.Note, error)
	GetAll(ctx context.Context, userId int) ([]domain.Note, error)
	Delete(ctx context.Context, userId, id int) error
	Update(ctx context.Context, userId, id int, newNote domain.UpdateNote) error
}

type Repository struct {
	Authorization
	Note
}

func GetRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: psql.GetAuthRepository(db),
		Note:          psql.GetNoteRepository(db),
	}
}
