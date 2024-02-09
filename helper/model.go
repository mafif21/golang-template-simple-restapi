package helper

import (
	"belajar-golang-restapi/model/domain"
	"belajar-golang-restapi/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
