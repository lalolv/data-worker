package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gookit/gcli/v2/progress"
)

// Build 编译输出json文件
func Build(outDir, fileName, fileExt string, count int, fields []interface{}) {
	// File full name
	filePath := fmt.Sprintf("./%s/%s.%s", outDir, fileName, fileExt)
	fmt.Println("Creating data file: ", filePath)
	// Check path exist
	if !CheckPathIsExist(outDir) {
		err := os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			fmt.Println("Create dir err:", err.Error())
		}
	}

	// Progressbar
	p := progress.Bar(count)
	p.Start()

	// Export
	switch fileExt {
	case "json":
		exportJSON(p, filePath, count, fields)
	case "csv":
		exportCSV(p, filePath, count, fields)
	}

	p.Finish()

}

func exportJSON(p *progress.Progress, filePath string, count int, fields []interface{}) {
	// rows
	rows := []map[string]interface{}{}
	// Add rows list
	for i := 0; i < count; i++ {
		// Generate a row data
		row := GenRowMap(i, fields)
		rows = append(rows, row)
		time.Sleep(time.Millisecond * SLEEP_DELAY)
		p.Advance()
	}

	// Build file in path
	b, err := json.Marshal(rows)
	if err != nil {
		fmt.Println("json err:", err)
	}
	err = ioutil.WriteFile(filePath, b, 0666)
	if err != nil {
		fmt.Println("Write json file err:", err.Error())
	}
}

func exportCSV(p *progress.Progress, filePath string, count int, fields []interface{}) {
	fp, err := os.Create(filePath) // 创建文件句柄
	if err != nil {
		fmt.Println("Create file err:", err.Error())
		return
	}
	defer fp.Close()
	// 写入UTF-8 BOM
	_, err = fp.WriteString("\xEF\xBB\xBF")
	if err != nil {
		fmt.Println("Write UTF-8 BOM err:", err.Error())
		return
	}
	// New writer
	w := csv.NewWriter(fp)

	// Add header
	headers := []string{}
	for _, v := range fields {
		field, ok := v.(map[string]interface{})
		if ok {
			headers = append(headers, field["name"].(string))
		}
	}
	err = w.Write(headers)
	if err != nil {
		fmt.Println("Write heades err:", err.Error())
		return
	}

	// Add data list
	for i := 0; i < count; i++ {
		// Generate a row data
		row := GenRowStrings(i, fields)
		err = w.Write(row)
		if err != nil {
			fmt.Println("Write row data err:", err.Error())
			break
		}
		time.Sleep(time.Millisecond * SLEEP_DELAY)
		p.Advance()
	}

	w.Flush()
}
