package main

import (
	"flag"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	DbPath string
}

func (appConfig *AppConfig) ParseArgs() {
	flag.StringVar(
		&appConfig.DbPath,
		"dbFile",
		withUserHomeDirectory("/Library/Containers/com.kryolokovlin.Noizio/Data/Library/Application Support/Noizio/Sounds.sqlite"),
		"Noizio Sounds.sqlite file",
	)

	flag.Parse()
}

func withUserHomeDirectory(pathToJoin string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can't establish $HOME directory. Fix your env or try `dbFile` flag.", err)
	}
	return path.Join(homeDir, pathToJoin)
}
