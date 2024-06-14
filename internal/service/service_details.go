package service

import (
	"context"
	"fmt"

	"records-app/internal/adapters/database"
	"records-app/internal/adapters/database/schemas"
	"records-app/settings"

	"github.com/bluele/gcache"
)

type ServiceRecordDetails struct {
	settings *settings.Settings
	cache    gcache.Cache
	db       database.AdapterMethods
}

func ServiceGetRecords(settings *settings.Settings, cache gcache.Cache, db database.AdapterMethods) *ServiceRecordDetails {
	return &ServiceRecordDetails{settings, cache, db}
}

func (s *ServiceRecordDetails) GetAllRecords(ctx context.Context) ([]schemas.Records, error) {
	return s.db.GetAllRecordsDB(ctx)
}

func (s *ServiceRecordDetails) GetByIDRecords(ctx context.Context, recordId int) (schemas.Records, error) {
	record, err := s.cache.Get(fmt.Sprintf("%d", recordId))
	if err == nil {
		return record.(schemas.Records), nil
	}

	record, err = s.db.GetByIDRecordsDB(ctx, recordId)
	if err != nil {
		return schemas.Records{}, err
	}
	s.cache.Set(interface{}(recordId), s.settings.TTL)
	return record.(schemas.Records), nil
}

func (s *ServiceRecordDetails) CreateRecords(ctx context.Context) (int, error) {
	recordId, err := s.db.CreateRecordsDB(ctx)
	if err != nil {
		return 0, err
	}
	s.cache.Set(interface{}(ctx), s.settings.TTL)
	return recordId, nil
}

func (s *ServiceRecordDetails) UpdateRecords(ctx context.Context, newId int,
	newTitle string, newDescription string) (schemas.Records, error) {
	// if !newId.IsValid() {
	// 	return errors.New("update structure has no values")
	// }
	s.cache.Set(interface{}(newId), s.settings.TTL)
	return s.db.UpdateRecordsDB(ctx, newId, newTitle, newDescription)
}

func (s *ServiceRecordDetails) DeleteRecords(ctx context.Context, recordId int) (int, error) {
	return s.db.DeleteRecordsDB(ctx, recordId)
}
