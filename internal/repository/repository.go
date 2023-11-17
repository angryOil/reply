package repository

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"reply/internal/domain"
	"reply/internal/repository/model"
	"reply/internal/repository/req"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

const (
	InternalServerError = "internal server error"
)

func (r Repository) Create(ctx context.Context, c req.Create) error {
	m := model.ToCreateModel(c)

	_, err := r.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		log.Println("Create NewInsert err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}

func (r Repository) Update(ctx context.Context, id int,
	validFunc func(domains []domain.Reply) (domain.Reply, error),
	updateFunc func(d domain.Reply) (req.Patch, error)) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Update BeginTx err: ", err)
		return errors.New(InternalServerError)
	}
	var models []model.Reply
	err = tx.NewSelect().Model(&models).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("Update NewSelect err: ", err)
		return errors.New(InternalServerError)
	}
	domains := model.ToDomainList(models)

	validDomain, err := validFunc(domains)
	if err != nil {
		return err
	}

	p, err := updateFunc(validDomain)
	if err != nil {
		return err
	}

	m := model.ToPatchModel(p)
	_, err = tx.NewInsert().Model(&m).On("conflict (id) do update").Exec(ctx)
	if err != nil {
		log.Println("Update NewInsert err: ", err)
		return errors.New(InternalServerError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Update tx.Commit err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}
