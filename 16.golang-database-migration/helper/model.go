package helper

import (
	"golang-restful-api/entity/domain"
	web "golang-restful-api/entity/web/category"
)

func ToCategoryResponse(c domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
