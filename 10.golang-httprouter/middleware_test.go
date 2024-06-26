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

type LogMiddleware struct {
	h http.Handler
}

type ErrorMiddleware struct {
	h http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive Request")
	middleware.h.ServeHTTP(w, r)
}

func (middleware *ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Error: ")
	defer func() {
		err := recover()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()

	middleware.h.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		panic("Ups")
	})

	logMiddleware := &LogMiddleware{h: r}
	
	errorMiddleware := &ErrorMiddleware{h: logMiddleware}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000/", nil)
	recorder := httptest.NewRecorder()

	errorMiddleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	assert.Equal(t, "Error : Ups", string(body))
}
