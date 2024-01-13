package service

import (
	"context"
	"errors"
	"fmt"

	config "github.com/Laem20957/records-app/configs"
	domain "github.com/Laem20957/records-app/internal/domains"
	repository "github.com/Laem20957/records-app/internal/repositories"
	"github.com/bluele/gcache"
)

type NoteService struct {
	config *config.Config
	cache  gcache.Cache
	repo   repository.Note
}

func GetNoteService(cfg *config.Config, cache gcache.Cache, repo repository.Note) *NoteService {
	return &NoteService{cfg, cache, repo}
}

func (s *NoteService) GetAll(ctx context.Context, userId int) ([]domain.Note, error) {
	return s.repo.GetAll(ctx, userId)
}

func (s *NoteService) GetByID(ctx context.Context, userId, id int) (domain.Note, error) {
	note, err := s.cache.Get(fmt.Sprintf("%d.%d", userId, id))
	if err == nil {
		return note.(domain.Note), nil
	}

	note, err = s.repo.GetByID(ctx, userId, id)
	if err != nil {
		return domain.Note{}, err
	}
	s.cache.Set(interface{}(userId), s.config.TTL)
	return note.(domain.Note), nil
}

func (s *NoteService) Create(ctx context.Context, userId int, note domain.Note) (int, error) {
	id, err := s.repo.Create(ctx, userId, note)
	if err != nil {
		return 0, err
	}
	s.cache.Set(interface{}(userId), s.config.TTL)
	return id, nil
}

func (s *NoteService) Update(ctx context.Context, userId, id int, newNote domain.UpdateNote) error {
	if !newNote.IsValid() {
		return errors.New("update structure has no values")
	}
	s.cache.Set(interface{}(newNote), s.config.TTL)
	return s.repo.Update(ctx, userId, id, newNote)
}

func (s *NoteService) Delete(ctx context.Context, userId, id int) error {
	return s.repo.Delete(ctx, userId, id)
}
