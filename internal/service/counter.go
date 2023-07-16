package service

import "github.com/begenov/tsarka-task/internal/repository"

type CouterService struct {
	repo repository.Counters
}

func NewCounterervice(repo repository.Counters) *CouterService {
	return &CouterService{
		repo: repo,
	}
}

func (s *CouterService) Add(key string, value int64) (int64, error) {
	return s.repo.Add(key, value)
}

func (s *CouterService) Sub(key string, value int64) (int64, error) {
	return s.repo.Sub(key, value)
}

func (s *CouterService) Get(key string) (int64, error) {
	return s.repo.Get(key)
}
