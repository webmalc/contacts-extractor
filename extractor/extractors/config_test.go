package extractors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/contacts-extractor/common/test"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.NotNil(t, c)
	assert.Equal(
		t, "https://app.yeahdesk.ru/api/clients/person/read", c.YeahdeskURL,
	)
	assert.Equal(
		t, "token", c.YeahdeskToken,
	)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
