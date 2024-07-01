package controller

import (
	webResponse "golang-restful-api/entity/web"
	webRequest "golang-restful-api/entity/web/category"
	"golang-restful-api/helper"
	"golang-restful-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryCreateRequest := webRequest.CategoryCreateRequest{}

	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Insert(r.Context(), categoryCreateRequest)

	webResponse := webResponse.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToJsonResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryUpdateRequest := webRequest.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	id :=
		helper.ConvertStringToInt(p.ByName("categoryId"))

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	webResponse := webResponse.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToJsonResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id :=
		helper.ConvertStringToInt(p.ByName("categoryId"))

	controller.CategoryService.Delete(r.Context(), id)

	webResponse := webResponse.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToJsonResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := helper.ConvertStringToInt(p.ByName("categoryId"))

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)

	webResponse := webResponse.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToJsonResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(r.Context())

	webResponse := webResponse.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToJsonResponse(w, webResponse)
}
