package core

import "fmt"

// ErrUnsupportedVersion is the error returned when the Noizio version is
// not supported
type ErrUnsupportedVersion struct{ Version Version }

func (e ErrUnsupportedVersion) Error() string {
	return fmt.Sprintf("Unsupported Noizio version: %s", e.Version)
}

// ErrNoVersionFound is the error returned when
// no `CFBundleShortVersionString` found in info.plist
type ErrNoVersionFound struct{ File string }

func (e ErrNoVersionFound) Error() string {
	return fmt.Sprintf("Can't find `CFBundleShortVersionString` in %s", e.File)
}
