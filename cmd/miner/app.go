package main

var (
	name    = "NoizioMiner ⚒"
	version = "0.1.0"
)

func main() {
	config := &AppConfig{}
	config.ParseArgs()
	app := App{config: config}
	app.Initialize()
	app.Run()
}
