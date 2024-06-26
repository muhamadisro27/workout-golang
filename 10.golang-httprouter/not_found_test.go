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

func TestNotFoundHandler(t *testing.T) {

	r := httprouter.New()

	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page Not Found")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	res := recorder.Result()

	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Page Not Found", string(body))
}
