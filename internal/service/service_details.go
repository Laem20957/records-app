package service

import (
	"context"
	"errors"
	"fmt"

	cfg "github.com/Laem20957/records-app/configurations"
	"github.com/Laem20957/records-app/internal/domain"
	repo "github.com/Laem20957/records-app/internal/repository"
	"github.com/bluele/gcache"
)

type ServiceRecordDetails struct {
	cfg   *cfg.Configurations
	cache gcache.Cache
	repo  repo.RepositoryMethods
}

func ServiceGetRecords(cfg *cfg.Configurations, cache gcache.Cache, repo repo.RepositoryMethods) *ServiceRecordDetails {
	return &ServiceRecordDetails{cfg, cache, repo}
}

func (srd *ServiceRecordDetails) GetAllRecords(ctx context.Context) ([]domain.Records, error) {
	return srd.repo.GetAllRecordsDB(ctx)
}

func (srd *ServiceRecordDetails) GetByIDRecords(ctx context.Context, userId, id int) (domain.Records, error) {
	record, err := srd.cache.Get(fmt.Sprintf("%d.%d", userId, id))
	if err == nil {
		return record.(domain.Records), nil
	}

	record, err = srd.repo.GetByIDRecordsDB(ctx, userId, id)
	if err != nil {
		return domain.Records{}, err
	}
	srd.cache.Set(interface{}(userId), srd.cfg.TTL)
	return record.(domain.Records), nil
}

func (srd *ServiceRecordDetails) CreateRecords(ctx context.Context, userId int, note domain.Records) (int, error) {
	id, err := srd.repo.CreateRecordsDB(ctx, userId, note)
	if err != nil {
		return 0, err
	}
	srd.cache.Set(interface{}(userId), srd.cfg.TTL)
	return id, nil
}

func (s *ServiceRecordDetails) UpdateRecords(ctx context.Context, userId, id int, record domain.UpdateRecord) error {
	if !record.IsValid() {
		return errors.New("update structure has no values")
	}
	s.cache.Set(interface{}(record), s.cfg.TTL)
	return s.repo.UpdateRecordsDB(ctx, userId, id, record)
}

func (s *ServiceRecordDetails) DeleteRecords(ctx context.Context, userId, id int) error {
	return s.repo.DeleteRecordsDB(ctx, userId, id)
}
