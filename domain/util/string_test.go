package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCaseWithSpace(t *testing.T) {
	assert.Equal(t, "SuperCode", ToCamelCase("Super Code"))
}

func TestCamelCaseWithoutSpace(t *testing.T) {
	assert.Equal(t, "Supercode", ToCamelCase("Supercode"))
}