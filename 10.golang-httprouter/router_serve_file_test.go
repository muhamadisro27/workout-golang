package golanghttprouter

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	r := httprouter.New()

	directory, _ := fs.Sub(resources, "resources")
	r.ServeFiles("/files/*filepath", http.FS(directory))

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/files/index.txt", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	assert.Equal(t, "Hello World", string(body))

}
