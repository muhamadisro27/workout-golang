package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (handler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	handler.Handler.ServeHTTP(w, r)
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After Execute Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(w, "Hello Middleware")
	})
	mux.HandleFunc(("/panic"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Executed")
		panic("Ups")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorMiddleware := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: errorMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
