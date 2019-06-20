package main

import core "github.com/TheMickeyMike/NoizioSampleMiner/pkg"

var (
	name    = "NoizioMiner ⚒"
	version = "0.1.0"
)

func main() {
	config := &core.AppConfig{}
	config.ParseArgs()
	app := App{config: config}
	app.Initialize()
	app.Run()
}
