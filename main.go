package main

import (
	"github.com/gookit/gcli/v2"
	"github.com/lalolv/data-worker/cmd"
)

func main() {
	// New app
	app := gcli.NewApp()
	app.Version = "0.1.1"
	app.Description = "Data worker,batch data generator."

	// Add command
	app.Add(cmd.BuildCommand())

	// Run
	app.Run()
}
