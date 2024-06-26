package golangjson

type Person struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Married   bool     `json:"is_married"`
	Hobbies   []string `json:"hobbies"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Street     string `json:"street"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
