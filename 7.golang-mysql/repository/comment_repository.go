package repository

import (
	"context"
	"golang-mysql/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, c entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
