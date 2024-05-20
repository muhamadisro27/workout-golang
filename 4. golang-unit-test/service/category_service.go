package service

import (
  "golang-unit-test/repository"
  "golang-unit-test/entity"
  "errors"
)

type CategoryService struct {
  Repository repository.CategoryRepository
}

func (s *CategoryService) Get(id string) (*entity.Category, error) {
  category := s.Repository.FindById(id)

  if category == nil {
    return nil, errors.New("Category Not Found")
  } else {
    return category, nil
  }
}