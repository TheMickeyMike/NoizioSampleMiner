package main

import "flag"

type AppConfig struct {
	DbPath          string
	SoundsDirectory string
}

func (appConfig *AppConfig) ParseArgs() {
	flag.StringVar(
		&appConfig.DbPath,
		"dbPath",
		"/Applications/Noizio.app/Contents/Resources/Sounds.sqlite",
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
