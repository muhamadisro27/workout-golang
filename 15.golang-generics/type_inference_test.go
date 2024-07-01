package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeInference(t *testing.T) {
	assert.Equal(t, int(100), Min(100, 100))
	assert.Equal(t, int64(100), Min(int64(100), int64(120)))
	assert.Equal(t, float64(100.20), Min(float64(100.20), float64(100.30)))
}
