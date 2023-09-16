package core

import (
	"context"
	"log/slog"

	"github.com/migregal/bmstu-iu7-ds-lab1/pkg/readiness"
)

type Core struct {
}

func New(lg *slog.Logger, probe *readiness.Probe) (*Core, error) {
	return &Core{}, nil
}

func (c *Core) AddPerson(ctx context.Context) error    {
	return nil
}

func (c *Core) GetPerson(ctx context.Context) error    {
	return nil
}

func (c *Core) GetPersons(ctx context.Context) error   {
	return nil
}

func (c *Core) UpdatePerson(ctx context.Context) error {
	return nil
}

func (c *Core) DeletePerson(ctx context.Context) error {
	return nil
}
