package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJson(t *testing.T) {
	person := Person{
		FirstName: "Muhamad",
		LastName:  "Isro",
		Age:       23,
		Married:   false,
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
