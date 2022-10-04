package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Must validate.
func TestVKValidator_Validate(t *testing.T) {
	v := NewVKValidator()

	assert.True(t, v.Validate("151680402"))
	assert.True(t, v.Validate("id25947743"))
	assert.True(t, v.Validate("ID25947743"))
	assert.True(t, v.Validate("206976714"))
	assert.True(t, v.Validate("4568934"))
	assert.True(t, v.Validate("45689"))

	assert.False(t, v.Validate("dddd"))
	assert.False(t, v.Validate("idid25947743"))
	assert.False(t, v.Validate("df23232323"))
	assert.False(t, v.Validate("df 23232323"))
	assert.False(t, v.Validate("df 23232323"))
	assert.False(t, v.Validate(""))
}

// Must return a new object.
func TestNewVKValidator(t *testing.T) {
	v := NewVKValidator()
	assert.NotNil(t, v)
}
