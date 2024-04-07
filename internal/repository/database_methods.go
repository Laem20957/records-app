package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Laem20957/records-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PSQL struct {
	DB *sqlx.DB
}

func RepositoryGetRecord(db *sqlx.DB) *PSQL {
	return &PSQL{DB: db}
}

func (psql *PSQL) CreateRecordsDB(ctx context.Context, userId int, record domain.Records) (int, error) {
	var id int

	createRecordQuery := fmt.Sprintf("INSERT INTO %s (uid, title, description) VALUES ($1, $2, $3) RETURNING id", recordsTable)
	row := psql.DB.QueryRowContext(ctx, createRecordQuery, userId, record.Title, record.Description)
	err := row.Scan(&id)
	return id, err
}

func (psql *PSQL) GetByIDRecordsDB(ctx context.Context, userId, id int) (domain.Records, error) {
	var record domain.Records

	query := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1 AND id = $2", recordsTable)
	err := psql.DB.GetContext(ctx, &record, query, userId, id)
	return record, err
}

func (psql *PSQL) GetAllRecordsDB(ctx context.Context) ([]domain.Records, error) {
	var records []domain.Records

	query := fmt.Sprintf("SELECT * FROM %s", recordsTable)
	err := psql.DB.SelectContext(ctx, &records, query)
	return records, err
}

func (psql *PSQL) DeleteRecordsDB(ctx context.Context, userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uid = $1 AND id = $2",
		recordsTable)

	_, err := psql.DB.ExecContext(ctx, query, userId, id)
	return err
}

func (psql *PSQL) UpdateRecordsDB(ctx context.Context, userId, id int, record domain.UpdateRecord) error {
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

	_, err := psql.DB.ExecContext(ctx, query, args...)
	return err
}
