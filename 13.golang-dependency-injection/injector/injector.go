//go:build wireinject
// +build wireinject

package injector

import (
	"golang-restful-api/app/database"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/router"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

// ProvideValidatorOptions returns the default options for the validator.
func ProvideValidatorOptions() []validator.Option {
	// Customize as needed
	return []validator.Option{}
}

// ProvideValidator returns a new validator.Validate instance.
func ProvideValidator(opts []validator.Option) *validator.Validate {
	return validator.New(opts...)
}

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:4000",
		Handler: middleware,
	}
}

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		database.GetConnection,
		ProvideValidatorOptions,
		ProvideValidator,
		categorySet,
		router.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
