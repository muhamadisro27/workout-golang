package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	// person := &Person{
	// 	FirstName: "Muhamad",
	// 	LastName:  "Isro Sabanur",
	// 	Age:       23,
	// 	Married:   false,
	// 	Hobbies:   []string{"Coding", "Gaming"},
	// 	Addresses: []Address{
	// 		{
	// 			Street:     "Jalan Kusuma",
	// 			Country:    "Indonesia",
	// 			PostalCode: "76142",
	// 		},
	// 		{
	// 			Street:     "Jalan Unocal",
	// 			Country:    "Indonesia",
	// 			PostalCode: "76141",
	// 		},
	// 	},
	// }

	p := new(Person)
	p.FirstName = "Muhamad"
	p.LastName = "Isro Sabanur"
	p.Age = 23
	p.Married = false
	p.Hobbies = []string{"Coding", "Gaming"}
	p.Addresses = []Address{
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
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"first_name":"Muhamad","last_name":"Isro Sabanur","age":23,"is_married":false,"hobbies":["Coding","Gaming"],"addresses":[{"street":"Jalan Kusuma","country":"Indonesia","postal_code":"76142"},{"street":"Jalan Unocal","country":"Indonesia","postal_code":"76141"}]}`

	jsonBytes := []byte(jsonString)

	person := new(Person)

	json.Unmarshal(jsonBytes, person)

	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
	fmt.Println(person.Age)
	fmt.Println(person.Married)
	fmt.Println(person.Hobbies)
	fmt.Println(person.Addresses)
}
