package golanggenerics

import (
	"fmt"
	"testing"
)

type TypeSetBag interface {
	string | int | int8 | int32 | int64 | float32 | float64
}

type Bag[T TypeSetBag] []T

func PrintBag[Z TypeSetBag](bag Bag[Z]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBag(t *testing.T) {
	numbers := Bag[int8]{1, 2, 3, 4, 5, 6, 7}

	PrintBag(numbers)

	names := Bag[string]{"Muhamad", "Isro", "Sabanur"}

	PrintBag(names)
}
