package lab2

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"bytes"
)

func TestCompute (t *testing.T) {
	input := "4 22 - 3 * 5 +"
	expected := "(4 - 22) * 3 + 5\n"
	rd := strings.NewReader(input)
	var buf bytes.Buffer
	handler := ComputeHandler {Input: rd, Output: &buf}
	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, expected, buf.String())
	}
} 