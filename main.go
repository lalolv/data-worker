package main

import (
	"github.com/gookit/gcli/v2"
	"github.com/lalolv/data-worker/cmd"
)

func main() {
	// New app
	app := gcli.NewApp()
	app.Version = "0.1.1"
	app.Description = "Data worker, data generation"

	// app.Add(cmd.ExampleCommand())
	app.Add(cmd.BuildCommand())

	// Run
	app.Run()
}
