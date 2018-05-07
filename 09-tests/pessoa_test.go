package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	pessoa := NewPessoa("Longa")

	assert.NotNil(t, pessoa)
	assert.Equal(t, "Nome: Longa", pessoa.String())
}
