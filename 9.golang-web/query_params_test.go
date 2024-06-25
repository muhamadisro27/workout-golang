package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleQueryParams(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:4000/hello?name=Isro&age=20", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Halo saya Isro dan saya berumur 20 tahun", string(body))
}

func MultipleQueryParams(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Halo saya %s dan saya berumur %s tahun", name, age)
	}
}

func TestMultipleParamsValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:4000/hello?name=Isro&name=Sabanur", nil)
	recorder := httptest.NewRecorder()

	MultipleParamsValues(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Isro,Sabanur", string(body))
}

func MultipleParamsValues(w http.ResponseWriter, r *http.Request) {
	var query url.Values = r.URL.Query()
	var names []string = query["name"]

	fmt.Fprint(w, strings.Join(names, ","))
}
