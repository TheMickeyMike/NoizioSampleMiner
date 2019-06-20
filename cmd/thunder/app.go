package main

var (
	name    = "NoizioThunder ⚡️"
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	config := &AppConfig{}
	config.ParseArgs()
	app := App{config: config}
	app.Initialize()
	app.Run()
}
