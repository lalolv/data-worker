package handler

import (
	"os"

	"github.com/lalolv/data-worker/workers"
)

// CheckPathIsExist 判断目录或文件是否存在  存在返回 true 不存在返回false
func CheckPathIsExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// GenRow Generate a row data
func GenRow(fields []interface{}) map[string]interface{} {
	// Set rand seed
	workers.Seed(0)

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
