package golanggenerics

import (
	"fmt"
	"testing"
)

type IData interface {
	string | int | int8 | int32 | int64 | float32 | float64
}

type Data[T IData] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first

	return first
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Muhamad",
		Second: "Isro",
	}

	fmt.Println(data.SayHello("Roozy"))

	fmt.Println(data.ChangeFirst("Roozy"))

	data2 := Data[int8]{
		First:  1,
		Second: 2,
	}

	fmt.Println(data, data2)
}
