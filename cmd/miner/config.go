package main

import (
	"flag"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

const infoPlistLocation = "/Applications/Noizio.app/Contents/Info.plist"

// AppConfig keeps app configuration
type AppConfig struct {
	DbPath            string
	SoundsDirectory   string
	InfoPlistLocation string
}

// ParseArgs map passed args to AppConfig
func (appConfig *AppConfig) ParseArgs() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can't establish $HOME directory. Please ensure $HOME env is set. Error: ", err)
	}

	flag.StringVar(
		&appConfig.DbPath,
		"dbPath",
		path.Join(homeDir, "/Library/Containers/com.kryolokovlin.Noizio/Data/Library/Application Support/Noizio/Sounds.sqlite"),
		"Noizio DB file",
	)

	flag.StringVar(
		&appConfig.SoundsDirectory,
		"soundsDirectory",
		"./sounds",
		"Destination directory for sounds files",
	)

	flag.Parse()

	appConfig.InfoPlistLocation = infoPlistLocation
}
