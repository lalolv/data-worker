package main

import (
	"flag"
	"fmt"

	"github.com/gookit/config"

	"github.com/lalolv/data-worker/handler"
)

func main() {
	// Get c flag
	cfg := flag.String("c", "", "the config")
	flag.Parse()
	fmt.Println(*cfg)

	// add driver for support yaml content
	// config.AddDriver(json.Driver)
	// Load json file
	err := config.LoadFiles(*cfg)
	if err != nil {
		panic(err)
	}

	// Get some basic params
	// format, _ := config.String("format")
	fileName, _ := config.String("file_name")
	filePath, _ := config.String("file_path")
	count, _ := config.Int("count")

	// Get all fields
	fields, _ := config.Get("fields")

	// Read dict
	userNames := handler.ReadLines([]int{3, 6, 10}, "./dict/usernames.txt")
	fmt.Println(userNames)

	// Build json file
	handler.Build(filePath, fileName, count, fields.([]interface{}))

}
