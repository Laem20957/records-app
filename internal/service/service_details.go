package service

import (
	"context"
	"errors"
	"fmt"

	"records-app/internal/adapters/database"
	"records-app/internal/adapters/database/schemas"
	"records-app/settings"

	"github.com/bluele/gcache"
)

type ServiceRecord struct {
	settings *settings.Settings
	cache    gcache.Cache
	db       database.AdapterMethods
}

func NewGetServiceRecord(settings *settings.Settings, cache gcache.Cache, db database.AdapterMethods) *ServiceRecord {
	return &ServiceRecord{settings, cache, db}
}

func (s *ServiceRecord) GetAllRecords(ctx context.Context) ([]schemas.Records, error) {
	if s == nil {
		return nil, errors.New("runtime error:" +
			"invalid memory address or nil pointer dereference")
	}
	return s.db.GetAllRecordsDB(ctx)
}

func (s *ServiceRecord) GetByIDRecords(ctx context.Context, recordId int) (schemas.Records, error) {
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

func (s *ServiceRecord) CreateRecords(ctx context.Context) (int, error) {
	recordId, err := s.db.CreateRecordsDB(ctx)
	if err != nil {
		return 0, err
	}
	s.cache.Set(interface{}(ctx), s.settings.TTL)
	return recordId, nil
}

func (s *ServiceRecord) UpdateRecords(ctx context.Context, newId int,
	newTitle, newDescription string) (schemas.Records, error) {
	// if !newId.IsValid() {
	// 	return errors.New("update structure has no values")
	// }
	s.cache.Set(interface{}(newId), s.settings.TTL)
	return s.db.UpdateRecordsDB(ctx, newId, newTitle, newDescription)
}

func (s *ServiceRecord) DeleteRecords(ctx context.Context, recordId int) (int, error) {
	return s.db.DeleteRecordsDB(ctx, recordId)
}
