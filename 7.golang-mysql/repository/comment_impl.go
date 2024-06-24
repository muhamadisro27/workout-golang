package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-mysql/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, c entity.Comment) (entity.Comment, error) {
	sqlExec := "insert into comments(email, comment) values(?,?)"

	result, err := repo.DB.ExecContext(ctx, sqlExec, c.Email, c.Comment)

	if err != nil {
		return c, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c, err
	}

	c.Id = int32(id)
	return c, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	sqlExec := "select id, email, comment from comments where id = ? LIMIT 1"

	rows, err := repo.DB.QueryContext(ctx, sqlExec, id)
	c := entity.Comment{}

	if err != nil {
		return c, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&c.Id, &c.Email, &c.Comment)
		return c, nil
	} else {
		return c, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlExec := "select id, email, comment from comments"

	rows, err := repo.DB.QueryContext(ctx, sqlExec)
	comments := []entity.Comment{}

	if err != nil {
		return comments, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
