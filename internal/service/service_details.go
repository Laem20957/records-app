package service

import (
	"context"
	"fmt"

	domain "records-app/internal/domain"
	repo "records-app/internal/repository"
	settings "records-app/settings"

	"github.com/bluele/gcache"
)

type ServiceRecordDetails struct {
	settings *settings.Settings
	cache    gcache.Cache
	repo     repo.RepositoryMethods
}

func ServiceGetRecords(settings *settings.Settings, cache gcache.Cache, repo repo.RepositoryMethods) *ServiceRecordDetails {
	return &ServiceRecordDetails{settings, cache, repo}
}

func (s *ServiceRecordDetails) GetAllRecords(ctx context.Context) ([]domain.Records, error) {
	return s.repo.GetAllRecordsDB(ctx)
}

func (s *ServiceRecordDetails) GetByIDRecords(ctx context.Context, recordId int) (domain.Records, error) {
	record, err := s.cache.Get(fmt.Sprintf("%d", recordId))
	if err == nil {
		return record.(domain.Records), nil
	}

	record, err = s.repo.GetByIDRecordsDB(ctx, recordId)
	if err != nil {
		return domain.Records{}, err
	}
	s.cache.Set(interface{}(recordId), s.settings.TTL)
	return record.(domain.Records), nil
}

func (s *ServiceRecordDetails) CreateRecords(ctx context.Context) (int, error) {
	recordId, err := s.repo.CreateRecordsDB(ctx)
	if err != nil {
		return 0, err
	}
	s.cache.Set(interface{}(ctx), s.settings.TTL)
	return recordId, nil
}

func (s *ServiceRecordDetails) UpdateRecords(ctx context.Context, newId int,
	newTitle string, newDescription string) (domain.Records, error) {
	// if !newId.IsValid() {
	// 	return errors.New("update structure has no values")
	// }
	s.cache.Set(interface{}(newId), s.settings.TTL)
	return s.repo.UpdateRecordsDB(ctx, newId, newTitle, newDescription)
}

func (s *ServiceRecordDetails) DeleteRecords(ctx context.Context, recordId int) (int, error) {
	return s.repo.DeleteRecordsDB(ctx, recordId)
}
