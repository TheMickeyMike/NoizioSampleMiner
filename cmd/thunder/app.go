package main


var (
	name    = "NoizioThunder ⚡️"
	version = "0.1.0"
)

func main() {
	config := &AppConfig{}
	config.ParseArgs()
	app := App{config: config}
	app.Initialize()
	app.Run()
}
