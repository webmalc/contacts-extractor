package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Extractor mocks the object.
type Extractor struct {
	mock.Mock
}

// Extract is method mock.
func (r *Extractor) Extract(
	sources []string, contacts []string,
) map[string][]string {
	arg := r.Called(sources, contacts)

	return arg.Get(0).(map[string][]string)
}
