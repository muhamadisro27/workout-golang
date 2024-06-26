package golanghttprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatternNamedP(t *testing.T) {
	r := httprouter.New()

	r.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		text := "Product " + p.ByName("id") + " Item " + p.ByName("itemId")

		fmt.Fprint(w, text)
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/products/1/items/1", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	response := rec.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 1", string(bytes))
}

func TestRouterCatchAllParams(t *testing.T) {
	r := httprouter.New()

	r.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Image " + p.ByName("image")
		fmt.Fprint(w, text)
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/images/small/profile.png", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	response := rec.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image /small/profile.png", string(bytes))
}
