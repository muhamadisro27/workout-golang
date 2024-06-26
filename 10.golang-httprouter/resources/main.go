package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Panic : " + i.(string)))
	}

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("Ups Elah")
	})

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
