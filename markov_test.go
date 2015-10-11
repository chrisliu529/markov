package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildIndice(t *testing.T) {
	result := BuildIndice("a b")
	assert.Equal(t, []string{"a"}, result["   "])

	result = BuildIndice("a b c")
	assert.Equal(t, []string{"c"}, result["a b"])

	result = BuildIndice("a b c a b d")
	assert.Equal(t, []string{"c", "d"}, result["a b"])
}
