package repository

import (
	"context"
	"fmt"

	"records-app/internal/domain"

	gorm "github.com/jinzhu/gorm"
)

var db, _ = ConnectDB()

type DatabaseServer struct {
	DB *gorm.DB
}

func RepositoryGetRecord(db *gorm.DB) *DatabaseServer {
	return &DatabaseServer{DB: db}
}

func (psql *DatabaseServer) GetAllRecordsDB(ctx context.Context) ([]domain.Records, error) {
	var records []domain.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Find(&records).Error; err != nil {
		logs.Error(err)
	}
	return records, nil
}

func (psql *DatabaseServer) GetByIDRecordsDB(ctx context.Context, recordId int) (domain.Records, error) {
	var record domain.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Where("id = ?", recordId).First(&record).Error; err != nil {
		logs.Error(err)
	}
	return record, nil
}

func (psql *DatabaseServer) CreateRecordsDB(ctx context.Context) (int, error) {
	var record domain.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Create(&record).Error; err != nil {
		logs.Error(err)
	}
	return record.ID, nil
}

func (psql *DatabaseServer) UpdateRecordsDB(ctx context.Context, newId int,
	newTitle string, newDescription string) (domain.Records, error) {

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Save(&domain.Records{
		ID: newId, Title: newTitle, Description: newDescription,
	}).Error; err != nil {
		logs.Error(err)
	}
	return domain.Records{
		ID: newId, Title: newTitle, Description: newDescription,
	}, nil
}

func (psql *DatabaseServer) DeleteRecordsDB(ctx context.Context, recordId int) (int, error) {
	var record domain.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Where("id = ?", recordId).Delete(&record).Error; err != nil {
		logs.Error(err)
	}
	return record.ID, nil
}
