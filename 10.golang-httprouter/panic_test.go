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

func TestPanic(t *testing.T) {
	r := httprouter.New()

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Panic : ", error)
	}

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("Ups")
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Panic : Ups", string(body))
}
