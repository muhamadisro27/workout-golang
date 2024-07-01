package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, int(1), ExperimentalMin(int(1), int(2)))
}
