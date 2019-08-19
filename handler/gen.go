package handler

import (
	"os"

	"github.com/lalolv/data-worker/workers"
	"github.com/lalolv/goutil"
)

// CheckPathIsExist 判断目录或文件是否存在  存在返回 true 不存在返回false
func CheckPathIsExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// GenRowMap Generate a row map data
func GenRowMap(index int, fields []interface{}) map[string]interface{} {
	row := map[string]interface{}{}
	for _, v := range fields {
		field, ok := v.(map[string]interface{})
		// append to row
		if ok {
			row[field["name"].(string)] = getCellVal(field, index)
		}
	}

	return row
}

// GenRowStrings Generate a string array
func GenRowStrings(index int, fields []interface{}) []string {
	row := []string{}
	for _, v := range fields {
		field, ok := v.(map[string]interface{})
		// append to row
		if ok {
			cellVal, _ := goutil.ToString(getCellVal(field, index))
			row = append(row, cellVal)
		}
	}

	return row
}

// Get a cell value
func getCellVal(field map[string]interface{}, index int) interface{} {
	val := ""
	// Use dict data
	if _, ok := field["dict"]; ok {
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

	return val
}
