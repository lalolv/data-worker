package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gookit/gcli/v2/progress"
)

// Build 编译输出json文件
func Build(filePath, fileName string, count int, fields []interface{}) {
	// rows
	rows := []map[string]interface{}{}
	// Progressbar
	p := progress.Bar(count)
	p.Start()
	// Add rows list
	for i := 0; i < count; i++ {
		// Generate a row data
		row := GenRow(i, fields)
		rows = append(rows, row)
		time.Sleep(time.Millisecond * 100)
		p.Advance()
	}
	p.Finish()

	// File full name
	file := fmt.Sprintf("./%s/%s.json", filePath, fileName)
	fmt.Println("Generate json file: ", file)
	// Check path exist
	if !CheckPathIsExist(filePath) {
		err := os.Mkdir(filePath, os.ModePerm)
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
