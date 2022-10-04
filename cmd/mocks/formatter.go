package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Formatter mocks the object.
type Formatter struct {
	mock.Mock
}

// Format is method mock.
func (r *Formatter) Format(
	results map[string][]string,
) string {
	arg := r.Called(results)

	return arg.Get(0).(string)
}
