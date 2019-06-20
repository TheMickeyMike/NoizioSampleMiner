package core

import (
	"flag"
	"log"
	"os"
	"path"
)

// AppConfig keeps app configuration
type AppConfig struct {
	DbPath          string
	SoundsDirectory string
}

// ParseArgs map passed args to AppConfig
func (appConfig *AppConfig) ParseArgs() {
	flag.StringVar(
		&appConfig.DbPath,
		"dbPath",
		withUserHomeDirectory("/Library/Containers/com.kryolokovlin.Noizio/Data/Library/Application Support/Noizio/Sounds.sqlite"),
		"Noizio DB file",
	)

	flag.StringVar(
		&appConfig.SoundsDirectory,
		"soundsDirectory",
		"./sounds",
		"Destination directory for sounds files",
	)
	flag.Parse()
}

func withUserHomeDirectory(pathToJoin string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can't establish $HOME directory. Fix your env or try `dbFile` flag. Error: ", err)
	}
	return path.Join(homeDir, pathToJoin)
}
