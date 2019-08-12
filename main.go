package main

import (
	"flag"

	"github.com/gookit/config"

	"github.com/lalolv/data-worker/handler"
	"github.com/lalolv/goutil"
)

func main() {
	// Get c flag
	cfg := flag.String("c", "", "the config")
	flag.Parse()

	// add driver for support yaml content
	// config.AddDriver(json.Driver)
	// Load json file
	err := config.LoadFiles(*cfg)
	if err != nil {
		panic(err)
	}

	// Get dict path
	dictPath, _ := config.String("dict_path")
	// Get build config
	buildInfo, _ := config.StringMap("build")
	// Get some basic params
	buildCount := goutil.Str2Int(buildInfo["count"])

	// Get all fields
	fields, _ := config.Get("fields")

	// Load dict data by fields
	handler.LoadDicts(fields.([]interface{}), float64(buildCount), dictPath)

	// Build json file
	handler.Build(buildInfo["path"], buildInfo["name"], buildCount, fields.([]interface{}))

}
