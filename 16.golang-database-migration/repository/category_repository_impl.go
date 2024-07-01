package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/entity/domain"
	"golang-restful-api/helper"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, c domain.Category) domain.Category {
	sql := "insert into category(name) values (?)"

	res, err := tx.ExecContext(ctx, sql, c.Name)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	c.Id = int(id)

	return c
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, c domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, c.Name, c.Id)
	helper.PanicIfError(err)

	return c
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, c domain.Category) {
	sql := "delete from category where id = ?"

	_, err := tx.ExecContext(ctx, sql, c.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	sql := "select id,name from category where id = ?"

	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicIfError(err)
	defer rows.Close()

	c := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&c.Id, &c.Name)
		helper.PanicIfError(err)
		return c, nil
	} else {
		return c, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "select id,name from category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
