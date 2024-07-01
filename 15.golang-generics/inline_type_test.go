package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NumberI interface {
	int | int64 | float64
}

func FindMin[T Number](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, int(1), FindMin(int(1), int(2)))
}

func FindInSlice[T interface{ []E }, E NumberI](values []E, target E) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func TestFindInSlice(t *testing.T) {
	assert.True(t, FindInSlice([]int64{1, 2, 3}, 3))
}

func GetFirst[T []E, E any](data T) E {
	if len(data) == 0 {
		return *new(E)
	}

	return data[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Muhamad", "Isro", "Sabanur",
	}

	first := GetFirst(names)

	assert.Equal(t, "Muhamad", first)
}
