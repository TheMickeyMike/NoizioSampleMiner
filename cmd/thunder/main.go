package main

import (
	"os"

	core "github.com/TheMickeyMike/NoizioThunderMiner/pkg"
	log "github.com/sirupsen/logrus"
)

// App is app backbone
type App struct {
	config *core.AppConfig
	store  *core.Store
}

// Initialize application state
func (app *App) Initialize() {
	log.Infof("%-13s: %s\n", "App name", name)
	log.Infof("%-13s: %s\n", "App version", version)
	log.Infoln()

	noizioVersion, err := core.GetNoizioVersion()
	if err != nil {
		log.Fatal("Can't determine Noizio version. Error: ", err)
	}
	if err := noizioVersion.IsSupported(); err != nil {
		log.Fatal(err)
	}

	app.store = core.NewStore(app.config.DbPath)
}

// Run application
func (app *App) Run() {
	defer app.store.Disconnect()

	log.Infof("Thundering sounds...")
	err := app.store.UpdateAllSounds()
	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Done. Peace Out 👊🏻")

	os.Exit(0)
}
