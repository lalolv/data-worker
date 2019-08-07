package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gookit/config"
	"github.com/gookit/gcli/v2/progress"

	"github.com/lalolv/data-worker/workers"
)

func main() {

	workers.Seed(0)
	num := workers.Number(1, 99)
	fmt.Println(num)

	// add driver for support yaml content
	// config.AddDriver(json.Driver)
	// Load json file
	err := config.LoadFiles("config/demo1.json")
	if err != nil {
		panic(err)
	}

	format, _ := config.String("format")
	fileName, _ := config.String("file_name")
	filePath, _ := config.String("file_path")
	count, _ := config.Int("count")

	fmt.Println(format, fileName, filePath, count)

	// Get all fields
	fields, _ := config.Get("fields")

	// rows
	rows := []map[string]interface{}{}
	// Progressbar
	p := progress.Bar(count)
	p.Start()
	// Add rows list
	for i := 0; i < count; i++ {
		row := genRow(fields.([]interface{}))
		rows = append(rows, row)
		time.Sleep(time.Millisecond * 100)
		p.Advance()
	}
	p.Finish()

	// File full name
	file := fmt.Sprintf("./%s/%s.json", filePath, fileName)
	fmt.Println("Generate json file: ", file)
	// Check path exist
	if !checkPathIsExist(filePath) {
		err = os.Mkdir(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("Create dir err:", err.Error())
		}
	}
	// Build file in path
	b, err := json.Marshal(rows)
	if err != nil {
		fmt.Println("json err:", err)
	}
	// fmt.Println(string(b))
	err = ioutil.WriteFile("./out/demo.json", b, 0666)
	if err != nil {
		fmt.Println("Write json file err:", err.Error())
	}
	fmt.Println("OK!")
}

/**
 * 判断目录或文件是否存在  存在返回 true 不存在返回false
 */
func checkPathIsExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
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
