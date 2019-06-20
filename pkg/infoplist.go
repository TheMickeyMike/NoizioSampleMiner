package core

import (
	"os"

	"howett.net/plist"
)

const infoPlistLocation = "/Applications/Noizio.app/Contents/Info.plist"

// InfoPlist represents info.plist structure
type InfoPlist struct {
	BundleShortVersion string `plist:"CFBundleShortVersionString"`
}

// GetNoizioVersion provides Noizio version
func GetNoizioVersion() (Version, error) {
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
	return Version(data.BundleShortVersion), nil
}
