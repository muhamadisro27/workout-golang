package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi Images")
	})

	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi Thumbnails")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi Page")
	})

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
