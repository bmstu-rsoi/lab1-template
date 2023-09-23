package service

import (
	"context"

	personModel "github.com/Astemirdum/person-service/internal/model"
)

func (s *Service) List(ctx context.Context) ([]personModel.Person, error) {
	return s.repo.List(ctx)
}

func (s *Service) Create(ctx context.Context, person personModel.Person) (int, error) {
	return s.repo.Create(ctx, person)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Get(ctx context.Context, id int) (personModel.Person, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) Update(ctx context.Context, person personModel.Person) (personModel.Person, error) {
	return s.repo.Update(ctx, person)
}
