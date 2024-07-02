package golangvalidation

import (
	"context"
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

func TestValidationTwoVaribale(t *testing.T) {
	var password string = "isro"
	var confirmPassword string = "isro"

	err := validate.VarWithValueCtx(context.Background(), password, confirmPassword, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleValidation(t *testing.T) {
	var user string = "isro167"

	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagValidation(t *testing.T) {
	user := "99"

	err := validate.Var(user, "required,number,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=6,max=10"`
	Password string `json:"password" validate:"required,alphanumunicode,min=5"`
}

func TestValidationStruct(t *testing.T) {
	user := LoginRequest{
		Username: "isroo167",
		Password: "isroo167",
	}

	err := validate.Struct(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	user := LoginRequest{
		Username: "isroo167",
		Password: "isroo167",
	}

	err := validate.Struct(&user)
	if err != nil {
		NewValidationError(err)
	}
}

// type RegisterRequest struct {
// 	Username        string `json:"username" validate:"required,alphanum,min=6,max=10"`
// 	Password        string `json:"password" validate:"required,alphanumunicode,min=5,eqfield=ConfirmPassword"`
// 	ConfirmPassword string `json:"confirm_password" validate:"required,min=5"`
// }

// func TestValidationCrossField(t *testing.T) {
// 	user := RegisterRequest{
// 		Username:        "isroo167",
// 		Password:        "isroo167",
// 		ConfirmPassword: "isroo1617",
// 	}

// 	err := validate.Struct(&user)

// 	if err != nil {
// 		NewValidationError(err)
// 	}
// }

type Address struct {
	City    string `validate:"required"`
	Country string `validate:"required"`
}

type School struct {
	Name string `validate:"required"`
}

type User struct {
	Id      string            `validate:"required"`
	Name    string            `validate:"required"`
	Address []Address         `validate:"required,dive"`
	Hobbies []string          `validate:"dive,required,min=3"`
	Schools map[string]School `validate:"required,dive,keys,required,endkeys"`
	Wallet  map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
}

func TestValidationNestedStruct(t *testing.T) {
	request := User{
		Id:   "roozy",
		Name: "roozy",
		Address: []Address{
			{
				City:    "Jakarta",
				Country: "Indonesia",
			},
			{
				City:    "Depok",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Coding",
		},
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationMap(t *testing.T) {
	request := User{
		Id:   "roozy",
		Name: "roozy",
		Address: []Address{
			{
				Country: "Jl. Sudirman",
				City:    "Jakarta",
			},
			{
				Country: "Jl. Margonda",
				City:    "Depok",
			},
		},
		Hobbies: []string{
			"Coding",
			"Reading",
			"Writing",
		},
		Schools: map[string]School{
			"SD": {
				Name: "Test School SD",
			},
			"SMP": {
				Name: "Test School SMP",
			},
			"SMA": {
				Name: "Test School SMA",
			},
		},
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationBasicMap(t *testing.T) {
	request := User{
		Id:   "roozy",
		Name: "roozy",
		Address: []Address{
			{
				Country: "Jl. Sudirman",
				City:    "Jakarta",
			},
			{
				Country: "Jl. Margonda",
				City:    "Depok",
			},
		},
		Hobbies: []string{
			"Coding",
			"Reading",
			"Writing",
		},
		Schools: map[string]School{
			"SD": {
				Name: "Test School SD",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "Test School SMA",
			},
		},
		Wallet: map[string]int{
			"GOPAY": 1000000,
			"BCA":   0,
			"":      1000,
		},
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAlias(t *testing.T) {
	newValidate := validator.New()

	newValidate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id   string `validate:"varchar"`
		Name string `validate:"varchar"`
	}

	seller := Seller{
		Id:   "1",
		Name: "",
	}

	err := newValidate.Struct(seller)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	user := LoginRequest{
		Username: "Isro167",
		Password: "test",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCustomPinValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type LoginRequest struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	user := LoginRequest{
		Phone: "085157708597",
		Pin:   "140010",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email|alphanum"`
		Password string `validate:"required"`
	}

	user := LoginRequest{
		Username: "isro12345",
		Password: "isro",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "isro@gmail.com",
		Email:    "isro@gmail.com",
		Phone:    "085157708597",
		Name:     "Muhamad Isro",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,numeric"`
	Password string `json:"password" validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {

	} else {
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {

	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "08123232132",
		Email:    "isro@sample.com",
		Phone:    "08123232132",
		Password: "rahasia",
	}

	err := validate.Struct(registerRequest)

	if err != nil {
		fmt.Println(err.Error())
	}
}
