package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is empty !")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hi %s", name)
	}
}

func TestResponseCodeSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000?name=Isro", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "Hi Isro", string(body))
}

func TestResponseCodeBadRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "Name is empty !", string(body))
}
