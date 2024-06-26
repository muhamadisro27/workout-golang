package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamJsonDecoder(t *testing.T) {
	reader, _ := os.Open("data.json")

	decoder := json.NewDecoder(reader)

	person := &Person{}
	_ = decoder.Decode(person)

	fmt.Println(person)

}

func TestStreamJsonEncoder(t *testing.T) {
	p := &Person{
		FirstName: "Muhamad",
		LastName:  "Isro Sabanur",
		Age:       23,
		Married:   false,
		Hobbies:   []string{"Coding", "Gaming", "Guitar"},
		Addresses: []Address{
			{
				Street:     "Jalan Kusuma",
				Country:    "Indonesia",
				PostalCode: "76142",
			},
			{
				Street:     "Jalan Unocal",
				Country:    "Indonesia",
				PostalCode: "76141",
			},
		},
	}

	w, _ := os.Create("data.json")

	encoder := json.NewEncoder(w)

	_ = encoder.Encode(p)

}
