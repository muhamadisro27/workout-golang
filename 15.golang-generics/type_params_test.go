package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneric(t *testing.T) {
	assert.True(t, true)

	var result string = Length[string]("Roozy")
	assert.Equal(t, result, "Roozy")

	var resultNumber int = Length[int](100)
	assert.Equal(t, resultNumber, 100)

}

func Length[T any](param T) T {
	return param
}
