package test

import (
	"os"
	"testing"

	"github.com/webmalc/contacts-extractor/common/config"
)

// Setups the tests.
func setUp() {
	os.Setenv("CE_ENV", "test")
	config.Setup()
}

// Run setups, runs and teardown the tests.
func Run(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
