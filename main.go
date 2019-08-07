package main

import (
	"fmt"
	"time"

	"github.com/gookit/config"
	"github.com/gookit/config/json"

	"github.com/lalolv/data-worker/workers"
)

func main() {

	workers.Seed(0)
	num := workers.Number(1, 99)
	fmt.Println(num)

	// add driver for support yaml content
	config.AddDriver(json.Driver)
	// Load json file
	err := config.LoadFiles("config/demo1.json")
	if err != nil {
		panic(err)
	}

	format, _ := config.String("format")
	path, _ := config.String("path")
	count, _ := config.Int("count")

	fmt.Println(format, path, count)

	// Get all fields
	fields, _ := config.Get("fields")

	// rows
	rows := []map[string]interface{}{}
	// Add rows list
	for i := 0; i < count; i++ {
		row := genRow(fields.([]interface{}))
		fmt.Println(row)
		rows = append(rows, row)
		time.Sleep(time.Millisecond * 100)
	}
}

// Generate a row data
func genRow(fields []interface{}) map[string]interface{} {
	row := map[string]interface{}{}
	for _, v := range fields {
		// Handle a field
		field, ok := v.(map[string]interface{})
		if ok {
			val := ""
			switch field["type"].(string) {
			case "string":
				// gen rand string
				val = workers.Letter()
			case "uuid":
				//
				val = workers.UUID()
			case "datetime":
				//
				val = workers.Date().Format("2006-01-02 15:04:05")
			}

			// append to row
			row[field["name"].(string)] = val
		}
	}

	return row
}
