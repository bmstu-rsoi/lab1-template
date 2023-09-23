package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Astemirdum/person-service/internal/errs"
	personModel "github.com/Astemirdum/person-service/internal/model"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

var qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (r *repository) List(ctx context.Context) ([]personModel.Person, error) {
	var pp []personModel.Person
	q := qb.Select("id", "name", "age", "address", "work").From(personsTableName)
	query, args, err := q.ToSql()
	if err != nil {
		return pp, err
	}
	err = r.db.SelectContext(ctx, &pp, query, args...)
	return pp, err
}

func (r *repository) Create(ctx context.Context, person personModel.Person) (int, error) {
	builder := qb.Insert(personsTableName).
		Columns("name", "age", "address", "work").
		Values(person.Name, person.Age, person.Address, person.Work).
		Suffix("returning id")
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	var id int
	if err := r.db.GetContext(ctx, &id, query, args...); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) Get(ctx context.Context, id int) (personModel.Person, error) {
	builder := qb.Select("id", "name", "age", "address", "work").
		From(personsTableName).
		Where(sq.Eq{
			"id": id,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return personModel.Person{}, err
	}
	r.log.Debug("person", zap.String("query", query), zap.Any("args", args))

	var u personModel.Person
	if err = r.db.GetContext(ctx, &u, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return personModel.Person{}, errs.ErrNotFound
		}
		return personModel.Person{}, err
	}
	return u, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	b := qb.Delete(personsTableName).Where(sq.Eq{"id": id})
	query, args, err := b.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *repository) Update(ctx context.Context, person personModel.Person) (personModel.Person, error) {
	builder := qb.Update(personsTableName)

	builder = builder.Set("name", person.Name)
	if person.Age != nil {
		builder = builder.Set("age", person.Age)
	}
	if person.Address != nil {
		builder = builder.Set("address", person.Address)
	}
	if person.Work != nil {
		builder = builder.Set("work", person.Work)
	}

	builder = builder.Where(sq.Eq{
		"id": person.ID,
	}).Suffix("returning *")

	query, args, err := builder.ToSql()
	if err != nil {
		return personModel.Person{}, err
	}

	if err = r.db.GetContext(ctx, &person, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return personModel.Person{}, errs.ErrNotFound
		}
		return personModel.Person{}, err
	}

	return person, nil
}
