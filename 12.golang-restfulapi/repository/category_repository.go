package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/entity/domain"
)

type CategoryRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, c domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, c domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, c domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
