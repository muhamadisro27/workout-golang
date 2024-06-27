package exception

import (
	"golang-restful-api/entity/web"
	"golang-restful-api/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	// if methodNotAllowed(w, r, err) {
	// 	return
	// }

	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

// func methodNotAllowed(w http.ResponseWriter, _ *http.Request, err interface{}) bool {
// 	w.WriteHeader(http.StatusInternalServerError)
// 	w.Header().Set("Content-Type", "application/json")

// 	webResponse := web.WebResponse{
// 		Code:   http.StatusInternalServerError,
// 		Status: "INTERNAL SERVER ERROR",
// 		Data:   err,
// 	}

// 	helper.WriteToJsonResponse(w, webResponse)
// }

func validationError(w http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToJsonResponse(w, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, _ *http.Request, err interface{}) bool {

	exception, ok := err.(NotFoundError)
	if ok {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToJsonResponse(w, webResponse)

		return true
	} else {
		return false
	}

}

func internalServerError(w http.ResponseWriter, _ *http.Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToJsonResponse(w, webResponse)
}
