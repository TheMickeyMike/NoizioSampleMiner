package core

import (
	"os"

	"howett.net/plist"
)

// InfoPlist represents info.plist structure
type InfoPlist struct {
	BundleShortVersion string `plist:"CFBundleShortVersionString"`
}

// GetNoizioVersion provides Noizio version
func GetNoizioVersion(infoPlistLocation string) (Version, error) {
	infoPlistFile, err := os.Open(infoPlistLocation)
	if err != nil {
		return "", err
	}
	defer infoPlistFile.Close()

	var data InfoPlist
	decoder := plist.NewDecoder(infoPlistFile)
	if err = decoder.Decode(&data); err != nil {
		return "", err
	}
	version := Version(data.BundleShortVersion)
	if version == "" {
		return version, ErrNoVersionFound{infoPlistLocation}
	}
	return version, nil
}
