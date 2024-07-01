package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int32 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, int(20), Min[int](20, 21))
	assert.Equal(t, int64(20), Min[int64](20, 21))
	assert.Equal(t, float32(20.0), Min[float32](20.0, 21.0))
	assert.Equal(t, float64(20.0), Min[float64](float64(20.0), float64(21.0)))
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(120)))
}
