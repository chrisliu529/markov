package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildIndice(t *testing.T) {
	result := BuildIndice("a b")
	assert.Equal(t, map[string][]string(nil), result)

	result = BuildIndice("a b c")
	assert.Equal(t, []string{"c"}, result["a b"])
}
