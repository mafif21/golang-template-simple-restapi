package service

import (
	"belajar-golang-restapi/exception"
	"belajar-golang-restapi/helper"
	"belajar-golang-restapi/model/domain"
	"belajar-golang-restapi/model/web"
	"belajar-golang-restapi/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	newCategory := service.Repository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(newCategory)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, request.Id) // menemukan kategori by id
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.Repository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId) // menemukan kategori by id
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.Repository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.Repository.FindAll(ctx, tx)
	var categoriesResponse []web.CategoryResponse

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, helper.ToCategoryResponse(category))
	}

	return categoriesResponse
}
