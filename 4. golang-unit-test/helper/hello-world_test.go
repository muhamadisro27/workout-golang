package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
  result := HelloWorld("Roozy")
	if result != "Hello Roozy" {
		panic("Result is not Hello Roozy")
	}
}