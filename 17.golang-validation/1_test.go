package golangvalidation

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz3(t *testing.T) {

	result := FizzBuzz(3)

	fmt.Println(result)
	assert.Equal(t, `["1","2","Buzz1"]`, result)
}

func TestFizzBuzz5(t *testing.T) {

	result := FizzBuzz(5)

	fmt.Println(result)
	assert.Equal(t, `["1","2","Buzz1","4","Buzz2"]`, result)
}

func TestFizzBuzz15(t *testing.T) {

	result := FizzBuzz(15)

	fmt.Println(result)
	assert.Equal(t, `["1","2","Buzz1","4","Buzz2","Buzz1","Buzz3","8","Buzz1","Buzz2","11","Buzz1","13","Buzz3","Fizz1"]`, result)
}

func FizzBuzz(iterate int) string {
	var result []string

	if iterate >= 10000 {
		return "Too many request"
	}

	for i := 1; i <= iterate; i++ {
		if i%3 == 0 && i%5 == 0 && i%7 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 && i%5 == 0 {
			result = append(result, "Fizz1")
		} else if i%3 == 0 && i%7 == 0 {
			result = append(result, "Fizz2")
		} else if i%5 == 0 && i%7 == 0 {
			result = append(result, "Fizz3")
		} else if i%3 == 0 {
			result = append(result, "Buzz1")
		} else if i%5 == 0 {
			result = append(result, "Buzz2")
		} else if i%7 == 0 {
			result = append(result, "Buzz3")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	byte, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(byte)
}
