package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/contacts-extractor/common/test"
)

func TestCSV_Format(t *testing.T) {
	csv := NewCSV()
	result := csv.Format(map[string][]string{
		"phone":  {"111", "222"},
		"emails": {"foo", "bar"},
	})
	assert.Contains(t, result, "111")
	assert.Contains(t, result, "222")
	assert.Contains(t, result, "foo")
	assert.Contains(t, result, "bar")
}

// Should create an object.
func TestNewCSV(t *testing.T) {
	csv := NewCSV()
	assert.NotNil(t, csv)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
