package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func TestValidation(t *testing.T) {
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {

	var user string = ""
	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}
