package repository

import (
	"context"

	personModel "github.com/Astemirdum/person-service/internal/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository interface {
	List(ctx context.Context) ([]personModel.Person, error)
	Create(ctx context.Context, person personModel.Person) (int, error)
	Get(ctx context.Context, id int) (personModel.Person, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, person personModel.Person) (personModel.Person, error)
}

type repository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewRepository(db *sqlx.DB, log *zap.Logger) (*repository, error) {
	return &repository{
		db:  db,
		log: log.Named("repo"),
	}, nil
}

const (
	personsTableName = `person.persons`
)
