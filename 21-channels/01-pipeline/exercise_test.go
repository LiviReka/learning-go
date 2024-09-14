package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := []float32{2, 4, 6, 8, 10, 12, 14, 16}
	assert.Equal(t, result, collector(multiplier(generator(input))))
}
