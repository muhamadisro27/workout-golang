package router

import (
	"fmt"
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	r := httprouter.New()

	r.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method Not Allowed")
	})

	r.GET("/api/categories", categoryController.FindAll)
	r.POST("/api/categories", categoryController.Insert)
	r.GET("/api/categories/:categoryId", categoryController.FindById)
	r.PUT("/api/categories/:categoryId", categoryController.Update)
	r.DELETE("/api/categories/:categoryId", categoryController.Delete)

	r.PanicHandler = exception.ErrorHandler

	return r
}
