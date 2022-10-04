package extractor

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Must merge the results.
func Test_merger_merge(t *testing.T) {
	m := newMerger(map[string]validator{})
	r := m.merge(map[string][]string{
		"vk":    {"id111", "id222"},
		"phone": {"111", "222", "222", "555"},
		"email": {"foo", "bar", "foo", "foo", "spam", "spam"},
	}, map[string][]string{
		"phone": {"111", "222", "222"},
		"email": {"foo", "bar", "foo", "foo"},
	})
	sort.Strings(r["vk"])
	sort.Strings(r["phone"])
	sort.Strings(r["email"])
	assert.Equal(t, []string{"id111", "id222"}, r["vk"])
	assert.Equal(t, []string{"111", "222", "555"}, r["phone"])
	assert.Equal(t, []string{"bar", "foo", "spam"}, r["email"])
}

// Must return a new object.
func Test_newMerger(t *testing.T) {
	assert.NotNil(t, newMerger(map[string]validator{}))
}
