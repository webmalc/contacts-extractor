package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/contacts-extractor/common/test"
)

// Must validate.
func TestEmailValidator_Validate(t *testing.T) {
	v := NewEmailValidator()

	assert.True(t, v.Validate("test@example.com"))
	assert.True(t, v.Validate("t@e.ru"))

	assert.False(t, v.Validate("ffff"))
	assert.False(t, v.Validate("ffff@ffff"))
	assert.False(t, v.Validate("@ffff.ru"))
	assert.False(t, v.Validate(""))
}

// Must return a new object.
func TestNewEmailValidator(t *testing.T) {
	v := NewEmailValidator()
	assert.NotNil(t, v)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
