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
func GenRow(index int, fields []interface{}) map[string]interface{} {
	row := map[string]interface{}{}
	for _, v := range fields {
		val := ""
		// Handle a field
		field, ok := v.(map[string]interface{})
		if ok {
			// Use dict data
			if _, ok = field["dict"]; ok {
				if dataColl, ok := dictData[field["name"].(string)]; ok {
					val = dataColl[index]
				}
			} else {
				// Set rand seed
				workers.Seed(0)
				// Get rand by type
				switch field["type"].(string) {
				case "string":
					// gen rand string
					val = workers.Letter()
				case "uuid":
					// uuid
					val = workers.UUID()
				case "datetime":
					// datatime
					val = workers.Date().Format("2006-01-02 15:04:05")
				}
			}

			// append to row
			row[field["name"].(string)] = val
		}
	}

	return row
}
