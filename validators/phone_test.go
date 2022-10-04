package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Must validate.
func TestPhoneValidator_Validate(t *testing.T) {
	v := NewPhoneValidator()

	assert.True(t, v.Validate("+79254445566"))
	assert.True(t, v.Validate("79254445566"))
	assert.True(t, v.Validate("+4953334532"))
	assert.True(t, v.Validate("495 3444455"))
	assert.True(t, v.Validate("+7(925)4445566"))
	assert.True(t, v.Validate("+7(925) 444-55-66"))
	assert.True(t, v.Validate("+7 (925) 444-55-66"))
	assert.True(t, v.Validate("+7 925 4445566"))
	assert.True(t, v.Validate("+7 925 444-55-66"))
	assert.True(t, v.Validate("+7 (902) 2235854"))

	assert.False(t, v.Validate("+79254445566ddddd"))
	assert.False(t, v.Validate("+79254445566 df df"))
	assert.False(t, v.Validate("+79254445566 "))
	assert.False(t, v.Validate("34-44"))
	assert.False(t, v.Validate(""))
}

// Must return a new object.
func TestNewPhoneValidator(t *testing.T) {
	v := NewPhoneValidator()
	assert.NotNil(t, v)
}
