package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/lalolv/data-worker/utils"
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
	val := field["value"].(string)

	for _, key := range parseFieldValue(field["value"].(string)) {
		var newVal string
		fullKey := fmt.Sprintf("%s.%s", field["name"].(string), key)
		if dataColl, ok := dictData[fullKey]; ok {
			newVal = dataColl[index]
		} else {
			// Set rand seed
			utils.Seed(0)
			// Get rand by type
			switch field["type"].(string) {
			case "string":
				// gen rand string
				newVal = workers.Letter()
			case "uuid":
				// uuid
				newVal = workers.UUID()
			case "mobile":
				// mobile
				newVal = workers.Mobile()
			case "idno":
				newVal = workers.IDNo()
			case "datetime":
				// datatime
				newVal = workers.Date().Format("2006-01-02 15:04:05")
			default:
				newVal = ""
			}
		}

		// 替换
		oldVal := fmt.Sprintf("{%s}", key)
		val = strings.Replace(val, oldVal, newVal, -1)
	}

	return val
}

// 解析出每个字段值的字典key
func parseFieldValue(val string) []string {
	keys := []string{}

	i := 0
	for {
		start := IndexStart(val, "{", i)
		if start <= 0 {
			break
		}
		end := IndexStart(val, "}", start)

		// 字段名称
		keys = append(keys, val[start+1:end])
		i = end
	}

	return keys
}
