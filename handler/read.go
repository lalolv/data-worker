package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/gookit/gcli/v2/progress"
	"github.com/lalolv/data-worker/workers"
	"github.com/lalolv/goutil"
)

// LoadDicts Load dict by fields
func LoadDicts(fields []interface{}, reCount float64, dictPath string) {
	dictData = map[string][]string{}
	// Read dict list file
	dictMaps := readDictList(dictPath)
	// for each field
	for _, v := range fields {
		val, ok := v.(map[string]interface{})
		if ok {
			if dict, ok := val["dict"]; ok {
				// Get dict info from list
				dictMap := dictMaps[dict.(string)].(map[string]interface{})
				// read data
				fieldsName := val["name"].(string)
				fmt.Println("Loading data for", fieldsName)
				// Read data
				count, _ := goutil.ToFloat64(dictMap["count"])
				dataColl := ReadLines(reCount, count, fmt.Sprintf("./%s/%s", dictPath, dictMap["file"]))
				// Rand data
				workers.ShuffleStrings(dataColl)
				// Add to dict
				dictData[fieldsName] = dataColl
			}
		}
	}
}

func readDictList(dictPath string) map[string]interface{} {
	listFile := fmt.Sprintf("./%s/.list.json", dictPath)
	// Check path exist
	if !CheckPathIsExist(listFile) {
		fmt.Println("No find dict list file!")
		return map[string]interface{}{}
	}

	data, err := ioutil.ReadFile(listFile)
	if err != nil {
		return map[string]interface{}{}
	}

	var dictMap map[string]interface{}
	json.Unmarshal(data, &dictMap)

	return dictMap
}

// ReadLines Read a few line
// @reCount return count
// @rowCount dict file row count
func ReadLines(reCount, rowCount float64, path string) []string {
	// open file
	ff, err := os.Open(path)
	if err != nil {
		fmt.Println("Read file err: ", err.Error())
		return []string{}
	}
	defer ff.Close()

	// a block count
	b0 := 0
	bn := int(math.Floor(rowCount / reCount))
	// rand index
	workers.Seed(0)
	index := workers.Number(b0, b0+bn)

	// scan line
	scanner := bufio.NewScanner(ff)

	// progress bar
	p := progress.Bar(int(rowCount))
	p.Start()

	// return list
	var list []string
	var line int
	for scanner.Scan() {
		lineText := scanner.Text()
		// append text to list
		if index == line {
			list = append(list, lineText)
			// break
			if int(reCount) == len(list) {
				break
			}
			// update index
			b0 += bn + 1
			// rand index
			workers.Seed(0)
			index = workers.Number(b0, b0+bn)
		}

		line++
		time.Sleep(time.Millisecond * 100)
		p.Advance()
	}

	p.Finish()

	return list
}
