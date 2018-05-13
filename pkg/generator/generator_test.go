package generator

import (
	"testing"

	"github.com/afief/ternakgopher/pkg/validator"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	cc := Generate("4")

	assert.Equal(t, len(cc), 19) //16 for the number, 3 for the dash -
	assert.True(t, validator.Validate(cc))
}
