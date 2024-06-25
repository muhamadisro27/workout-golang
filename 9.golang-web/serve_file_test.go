package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/404.html
var resource404 string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Content-Disposition", "attachment; filename=test.html")
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceOK)
	} else {
		fmt.Fprint(w, resource404)
	}
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/404.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:4000",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(w, r, "./resources/"+file)
}

func TestDownload(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:4000",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
