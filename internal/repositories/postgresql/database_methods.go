package postgresql

import (
	"context"
	"fmt"
	"strings"

	domain "github.com/Laem20957/records-app/internal/domains"
	"github.com/jmoiron/sqlx"
)

type PSQL struct {
	dblib *sqlx.DB
}

func RepositoryGetRecord(dblib *sqlx.DB) *PSQL {
	return &PSQL{dblib: dblib}
}

func (psql *PSQL) CreateRecords(ctx context.Context, userId int, record domain.Record) (int, error) {
	var id int

	createRecordQuery := fmt.Sprintf("INSERT INTO %s (uid, title, description) VALUES ($1, $2, $3) RETURNING id",
		recordsTable)
	row := psql.dblib.QueryRowContext(ctx, createRecordQuery, userId, record.Title, record.Description)
	err := row.Scan(&id)

	return id, err
}

func (psql *PSQL) GetByIDRecords(ctx context.Context, userId, id int) (domain.Record, error) {
	var record domain.Record

	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE uid = $1 AND id = $2",
		recordsTable)
	err := psql.dblib.GetContext(ctx, &record, query, userId, id)

	return record, err
}

func (psql *PSQL) GetAllRecords(ctx context.Context, userId int) ([]domain.Record, error) {
	var records []domain.Record

	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE uid = $1",
		recordsTable)
	err := psql.dblib.SelectContext(ctx, &records, query, userId)

	return records, err
}

func (psql *PSQL) DeleteRecords(ctx context.Context, userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uid = $1 AND id = $2",
		recordsTable)

	_, err := psql.dblib.ExecContext(ctx, query, userId, id)

	return err
}

func (psql *PSQL) UpdateRecords(ctx context.Context, userId, id int, record domain.UpdateRecord) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if record.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *record.Title)
		argId++
	}

	if record.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *record.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND uid=$%d", recordsTable, setQuery, argId, argId+1)
	args = append(args, id, userId)

	_, err := psql.dblib.ExecContext(ctx, query, args...)

	return err
}
