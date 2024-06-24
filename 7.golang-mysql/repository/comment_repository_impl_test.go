package repository

import (
	"context"
	"fmt"
	belajar_golang_db "golang-mysql"
	"golang-mysql/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_db.GetConnection())

	ctx := context.Background()

	c := entity.Comment{
		Email:   "isro@gmail.com",
		Comment: "Test Repository",
	}

	c, err := commentRepository.Insert(ctx, c)

	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_db.GetConnection())

	ctx := context.Background()

	c, err := commentRepository.FindById(ctx, 43)

	if err != nil {
		panic(err)
	}

	if assert.NotNil(t, c) {
		assert.Equal(t, c.Id, int32(43))
		assert.Equal(t, c.Email, "isro@gmail.com")
		assert.Equal(t, c.Comment, "Test Repository")
	}

}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_db.GetConnection())

	ctx := context.Background()

	c, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	if assert.NotNil(t, c) {
		assert.Equal(t, len(c), 33)
	}
}
