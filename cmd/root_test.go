package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/contacts-extractor/cmd/mocks"
	"github.com/webmalc/contacts-extractor/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	e := &mocks.Extractor{}
	f := &mocks.Formatter{}
	cr := NewCommandRouter(m, e, f)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.ErrorLogger{}
	e := &mocks.Extractor{}
	f := &mocks.Formatter{}
	cr := NewCommandRouter(m, e, f)
	assert.Equal(t, m, cr.logger)
	assert.Equal(t, e, cr.extractor)
	assert.Equal(t, f, cr.formatter)
	assert.NotNil(t, cr.rootCmd)
}

// Should run the extract command.
func TestCommandRouter_extract(t *testing.T) {
	m := &mocks.ErrorLogger{}
	e := &mocks.Extractor{}
	f := &mocks.Formatter{}
	cr := NewCommandRouter(m, e, f)
	command := &cobra.Command{}
	command.Flags().StringSlice("sources", []string{}, "")
	command.Flags().StringSlice("contacts", []string{}, "")

	e.On(
		"Extract", mock.Anything, mock.Anything,
	).Return(map[string][]string{}).Once()
	f.On("Format", mock.Anything).Return("").Once()
	cr.extract(command, []string{})
	e.AssertExpectations(t)
	f.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
