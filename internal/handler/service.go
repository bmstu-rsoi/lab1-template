package handler

import (
	"context"

	personModel "github.com/Astemirdum/person-service/internal/model"
	"github.com/Astemirdum/person-service/internal/service"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type PersonService interface {
	List(ctx context.Context) ([]personModel.Person, error)
	Create(ctx context.Context, person personModel.Person) error
	Get(ctx context.Context, id int) (personModel.Person, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, person personModel.Person) (personModel.Person, error)
}

var _ PersonService = (*service.Service)(nil)
