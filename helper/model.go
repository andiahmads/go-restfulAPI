package helper

import (
	"github.com/andiahmads/go-restfulAPI/model/domain"
	"github.com/andiahmads/go-restfulAPI/model/web"
)

func ToCreateCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponse []web.CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCreateCategoryResponse(category))
	}
	return categoryResponse
}
