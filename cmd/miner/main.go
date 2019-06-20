package main

import (
	"fmt"
	"os"

	core "github.com/TheMickeyMike/NoizioThunderMiner/pkg"
	log "github.com/sirupsen/logrus"
)

// App is app backbone
type App struct {
	config      *AppConfig
	store       *core.Store
	fileManager *core.FileManager
}

// Initialize application state
func (app *App) Initialize() {
	log.Infof("%-13s: %s\n", "App name", name)
	log.Infof("%-13s: %s (%s)\n", "App version", version, commit)
	log.Infoln()

	noizioVersion, err := core.GetNoizioVersion(app.config.InfoPlistLocation)
	if err != nil {
		log.Fatal("Can't determine Noizio version. Error: ", err)
	}
	if err := noizioVersion.IsSupported(); err != nil {
		log.Fatal(err)
	}

	app.store = core.NewStore(app.config.DbPath)
	app.fileManager = core.NewFileManager(app.config.SoundsDirectory)
}

// Run application
func (app *App) Run() {
	defer app.store.Disconnect()

	sounds, err := app.store.GetAllSounds()
	if err != nil {
		log.Fatal(err)
	}

	if err := core.EnsureDirectoryExist(app.config.SoundsDirectory); err != nil {
		log.Fatal("Can't create sounds directory: Error: ", err)
	}

	log.Infof("Saving sounds to directory: %s\n", app.config.SoundsDirectory)

	for _, sound := range sounds {
		app.fileManager.SaveToFile(fmt.Sprintf("%s.caf", sound.Title()), sound.Data())
	}

	log.Infoln("Done. Peace Out 👊🏻")

	os.Exit(0)
}
