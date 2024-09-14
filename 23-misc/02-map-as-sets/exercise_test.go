package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestContain(t *testing.T) {
	var stdin bytes.Buffer
	stdin.WriteString("One Mind there is; but under it two principles contend. The Mind lets in the light, then the dark, in interaction; so time is generated. At the end Mind awards victory to the light; time ceases and the Mind is complete. He causes things to look different so it would appear time has passed. Matter is plastic in the face of Mind.")

	assert.Equal(t, true, contain(&stdin, "mind"), "The result should be: true")
}
