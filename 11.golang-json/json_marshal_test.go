package golangjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Type of "+string(bytes), reflect.TypeOf(string(bytes)))
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Isro")
	logJson(1)
	logJson(true)
	logJson(map[string]string{
		"New York": "New",
		"Test":     "Test Aja",
	})
	logJson([]string{"Muhamad", "Isro", "Sabanur"})
}
