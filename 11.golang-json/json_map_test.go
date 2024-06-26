package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMap(t *testing.T) {
	jsonRequest := `{"id":"12345","name":"Apple M1","price":10000000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}

	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
