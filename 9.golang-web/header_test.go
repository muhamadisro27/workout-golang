package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const URL string = "https://enwym7shc9geo.x.pipedream.net"

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	assert.Equal(t, "application/json", string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Golang")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "OK", string(body))
	assert.Equal(t, "OK", string(body))
	assert.Equal(t, "Golang", response.Header.Get("X-Powered-By"))
}