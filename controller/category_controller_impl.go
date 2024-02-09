package controller

import (
	"belajar-golang-restapi/helper"
	"belajar-golang-restapi/model/web"
	"belajar-golang-restapi/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService}
}

func (category CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryRequest)

	categoryResponse := category.CategoryService.Create(request.Context(), categoryRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "new category has been created",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (category CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	updateCategoryResponse := category.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "category has been updated",
		Data:   updateCategoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (category CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	category.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "category has been deleted",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (category CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	foundCategory := category.CategoryService.FindById(request.Context(), id)
	var webResponse = web.WebResponse{
		Code:   http.StatusOK,
		Status: "get category by id " + strconv.Itoa(id),
		Data:   foundCategory,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (category CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categories := category.CategoryService.FindAll(request.Context())
	var webResponse = web.WebResponse{
		Code:   http.StatusOK,
		Status: "get all categories",
		Data:   categories,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
