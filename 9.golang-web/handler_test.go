package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}
	server := http.Server{
		Addr:    "localhost:4000",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.URL)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.Host)
		fmt.Fprintln(w, r.RemoteAddr)
	}
	server := http.Server{
		Addr:    "localhost:4000",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:4000/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello World", string(body))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestSayHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:4000/hello?name=Isro", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello Isro", string(body))
}
