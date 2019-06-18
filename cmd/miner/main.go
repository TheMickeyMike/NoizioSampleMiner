package main

import (
	"os"

	core "github.com/TheMickeyMike/NoizioSampleMiner/pkg"
	log "github.com/sirupsen/logrus"
)

type App struct {
	config      *AppConfig
	store       *core.Store
	fileManager *core.FileManager
}

func (app *App) Initialize() {
	log.Printf("%-13s: %s\n", "App name", name)
	log.Printf("%-13s: %s\n", "App version", version)

	app.store = core.NewStore(app.config.DbPath)
	app.fileManager = core.NewFileManager(app.config.SoundsDirectory)
}

func (app *App) Run() {
	defer app.store.Disconnect()

	log.Infof("Saving sounds to directory: %s\n", app.config.SoundsDirectory)
	for _, sound := range app.store.GetAllSounds() {
		app.fileManager.SaveToFile(sound.Title(), sound.Data())
	}
	log.Infoln("Done. Peace Out 👊🏻")

	os.Exit(0)
}
