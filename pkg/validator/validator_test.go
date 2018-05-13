package validator

import (
	"testing"

	"github.com/afief/ternakgopher/pkg/generator"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	cc := generator.Generate("4")

	assert.True(t, Validate(cc))
}
