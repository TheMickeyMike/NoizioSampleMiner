package main

import (
	"os"

	core "github.com/TheMickeyMike/NoizioSampleMiner/pkg"
	log "github.com/sirupsen/logrus"
)

type App struct {
	config *AppConfig
	store  *core.Store
}

func (app *App) Initialize() {
	log.Infof("%-13s: %s\n", "App name", name)
	log.Infof("%-13s: %s\n", "App version", version)
	log.Infoln()

	core.Load()

	app.store = core.NewStore(app.config.DbPath)
}

func (app *App) Run() {
	defer app.store.Disconnect()

	err := app.store.UpdateAllSounds()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Success!")
	os.Exit(0)
}
