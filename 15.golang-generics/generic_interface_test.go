package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

func TestGenericInterface(t *testing.T) {
	data := MyData[string]{}
	result := ChangeValue(&data, "Roozy")

	assert.Equal(t, "Roozy", result)
	assert.Equal(t, "Roozy", data.GetValue())
}
