package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   *string `json:"last_name"`
}

func TestDecodeJSON(t *testing.T) {

	jsonRequest := `{"first_name":"Muhamad","middle_name":"Isro","last_name":"Sabanur"}`
	jsonBytes := []byte(jsonRequest)

	jsonRequest2 := `{"first_name":"Muhamad","middle_name":"Isro"}`
	jsonBytes2 := []byte(jsonRequest2)

	customer := &Customer{}
	customer2 := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonBytes2, customer2)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName, customer.MiddleName, *customer.LastName)
	fmt.Println(customer2.FirstName, customer2.MiddleName)
	if customer2.LastName != nil {
		fmt.Println(*customer2.LastName)
	} else {
		fmt.Println("Not Set")
	}
}
