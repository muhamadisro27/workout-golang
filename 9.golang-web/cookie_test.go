package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {

	cookie := new(http.Cookie)

	cookie.Name = "X-PZN-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success create cookie")
}

func TestSetCookie(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000?name=Isro", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		assert.Equal(t, "Isro", cookie.Value)
		assert.Equal(t, "/", cookie.Path)
		assert.Equal(t, "X-PZN-Name", cookie.Name)
	}
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		fmt.Fprintf(w, "Hello %s", cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:4000", nil)
	recorder := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Isro"
	request.AddCookie(cookie)

	GetCookie(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello Isro", string(body))
}
