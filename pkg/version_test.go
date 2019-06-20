package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Check if error returned when unsupported version passed
func TestIsSupportedWithUnsupportedVersion(t *testing.T) {
	expectedErrorMsg := "Unsupported Noizio version: 1.0"
	version := Version("1.0")
	err := version.IsSupported()
	assert.EqualError(t, err, expectedErrorMsg)
}

// No error returned when supported version passed
func TestIsSupportedWithSupportedVersion(t *testing.T) {
	version := Version("2.0")
	err := version.IsSupported()
	assert.NoError(t, err)
}
