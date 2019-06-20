package core

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const infoPlistSnippet = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleShortVersionString</key>
	<string>%s</string>
</dict>
</plist>
`

func TestGetNoizioVersionWhenNoFile(t *testing.T) {
	_, err := GetNoizioVersion("/tmp")
	assert.Error(t, err)
}

func TestGetNoizioVersionWhenNoVersionInFile(t *testing.T) {
	file, err := ioutil.TempFile("", "info.plist")
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	expectedErrorMsg := "Can't find `CFBundleShortVersionString` in " + file.Name()

	_, err = GetNoizioVersion(file.Name())
	assert.EqualError(t, err, expectedErrorMsg)

}

func TestGetNoizioVersion(t *testing.T) {
	expectedVersion := Version("2.0.1")

	file, err := ioutil.TempFile("", "info.plist")
	if err != nil {
		t.Fatal(err)
	}

	file.WriteString(fmt.Sprintf(infoPlistSnippet, expectedVersion))

	file.Close()

	version, err := GetNoizioVersion(file.Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedVersion, version)
}
