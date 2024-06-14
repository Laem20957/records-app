package database

import (
	"context"
	"fmt"

	"records-app/internal/adapters/database/schemas"

	"github.com/jinzhu/gorm"
)

var db, _ = ConnectDB()

type DatabaseRecordORM struct {
	DB *gorm.DB
}

func AdapterGetRecord(db *gorm.DB) *DatabaseRecordORM {
	return &DatabaseRecordORM{DB: db}
}

func (d *DatabaseRecordORM) GetAllRecordsDB(ctx context.Context) ([]schemas.Records, error) {
	var records []schemas.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Find(&records).Error; err != nil {
		logs.Error(err)
	}
	return records, nil
}

func (d *DatabaseRecordORM) GetByIDRecordsDB(ctx context.Context, recordId int) (schemas.Records, error) {
	var record schemas.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Where("id = ?", recordId).First(&record).Error; err != nil {
		logs.Error(err)
	}
	return record, nil
}

func (d *DatabaseRecordORM) CreateRecordsDB(ctx context.Context) (int, error) {
	var record schemas.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Create(&record).Error; err != nil {
		logs.Error(err)
	}
	return record.ID, nil
}

func (d *DatabaseRecordORM) UpdateRecordsDB(ctx context.Context, newId int,
	newTitle string, newDescription string) (schemas.Records, error) {

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Save(&schemas.Records{
		ID: newId, Title: newTitle, Description: newDescription,
	}).Error; err != nil {
		logs.Error(err)
	}
	return schemas.Records{
		ID: newId, Title: newTitle, Description: newDescription,
	}, nil
}

func (d *DatabaseRecordORM) DeleteRecordsDB(ctx context.Context, recordId int) (int, error) {
	var record schemas.Records

	if err := db.Table(fmt.Sprintf("records_app.%s", recordsTable)).Where("id = ?", recordId).Delete(&record).Error; err != nil {
		logs.Error(err)
	}
	return record.ID, nil
}
