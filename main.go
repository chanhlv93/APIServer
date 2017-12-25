package main

import (
	"github.com/chanhlv93/APIServer/app"
	"github.com/chanhlv93/APIServer/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
