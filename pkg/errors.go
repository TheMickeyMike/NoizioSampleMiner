package core

import "fmt"

// ErrUnsupportedVersion is the error returned when the Noizio version is
// not supported
type ErrUnsupportedVersion struct{ Version Version }

func (e ErrUnsupportedVersion) Error() string {
	return fmt.Sprintf("Unsupported Noizio version: %s", e.Version)
}
