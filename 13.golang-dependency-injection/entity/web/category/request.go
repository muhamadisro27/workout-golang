package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=100,min=1"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=100,min=1"`
}
