package postgresql

import (
	"context"
	"fmt"

	"github.com/Laem20957/records-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PSQLAUTH struct {
	dblib *sqlx.DB
}

func RepositoryGetAuth(db *sqlx.DB) *PSQLAUTH {
	return &PSQLAUTH{dblib: db}
}

func (repo *PSQLAUTH) CreateUserDB(ctx context.Context, user domain.Users) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	var id int
	row := repo.dblib.QueryRowContext(ctx, query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *PSQLAUTH) CreateTokenDB(ctx context.Context, token domain.RefreshSession) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, token, expires_at) values ($1, $2, $3)",
		refreshTokensTable)
	_, err := repo.dblib.ExecContext(ctx, query, token.UserID, token.Token, token.ExpiresAt)
	return err
}

func (repo *PSQLAUTH) GetUserDB(ctx context.Context, username, password string) (domain.Users, error) {
	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 AND password_hash=$2", usersTable)

	var user domain.Users
	err := repo.dblib.GetContext(ctx, &user, query, username, password)
	return user, err
}

func (repo *PSQLAUTH) GetTokenDB(ctx context.Context, token string) (domain.RefreshSession, error) {
	query := fmt.Sprintf("SELECT id, user_id, token, expires_at FROM %s WHERE token = $1",
		refreshTokensTable)

	var t domain.RefreshSession
	if err := repo.dblib.GetContext(ctx, &t, query, token); err != nil {
		return t, err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", refreshTokensTable)
	_, err := repo.dblib.ExecContext(ctx, query, t.UserID)
	return t, err
}
