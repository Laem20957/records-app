package service

import (
	"context"
	"errors"
	"fmt"

	config "github.com/Laem20957/records-app/configuration"
	"github.com/Laem20957/records-app/internal/domain"
	"github.com/Laem20957/records-app/internal/repository"
	"github.com/bluele/gcache"
)

type ServiceRecordDetails struct {
	config *config.Config
	cache  gcache.Cache
	repo   repository.RepositoryMethods
}

func ServiceGetRecords(cfg *config.Config, cache gcache.Cache, repo repository.RepositoryMethods) *ServiceRecordDetails {
	return &ServiceRecordDetails{cfg, cache, repo}
}

func (srd *ServiceRecordDetails) GetAllRecords(ctx context.Context, userId int) ([]domain.Record, error) {
	return srd.repo.GetAllRecordsDB(ctx, userId)
}

func (srd *ServiceRecordDetails) GetByIDRecords(ctx context.Context, userId, id int) (domain.Record, error) {
	record, err := srd.cache.Get(fmt.Sprintf("%d.%d", userId, id))
	if err == nil {
		return record.(domain.Record), nil
	}

	record, err = srd.repo.GetByIDRecordsDB(ctx, userId, id)
	if err != nil {
		return domain.Record{}, err
	}
	srd.cache.Set(interface{}(userId), srd.config.TTL)
	return record.(domain.Record), nil
}

func (srd *ServiceRecordDetails) CreateRecords(ctx context.Context, userId int, note domain.Record) (int, error) {
	id, err := srd.repo.CreateRecordsDB(ctx, userId, note)
	if err != nil {
		return 0, err
	}
	srd.cache.Set(interface{}(userId), srd.config.TTL)
	return id, nil
}

func (s *ServiceRecordDetails) UpdateRecords(ctx context.Context, userId, id int, newNote domain.UpdateRecord) error {
	if !newNote.IsValid() {
		return errors.New("update structure has no values")
	}
	s.cache.Set(interface{}(newNote), s.config.TTL)
	return s.repo.UpdateRecordsDB(ctx, userId, id, newNote)
}

func (s *ServiceRecordDetails) DeleteRecords(ctx context.Context, userId, id int) error {
	return s.repo.DeleteRecordsDB(ctx, userId, id)
}
