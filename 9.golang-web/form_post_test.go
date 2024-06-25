package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Isro&lastName=Sabanur")

	request := httptest.NewRequest("POST", "http://localhost:4000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello Isro Sabanur", string(body))
}