package handler

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gookit/gcli/v2/progress"
	"github.com/lalolv/data-worker/workers"
)

// LoadDicts Load dict by fields
func LoadDicts(fields []interface{}, reCount float64, dictPath string) {
	dictData = map[string][]string{}
	// for each field
	for _, v := range fields {
		val, ok := v.(map[string]interface{})
		if ok {
			if dict, ok := val["dict"]; ok {
				fmt.Println("dict", dict)
				// read data
				fieldsName := val["name"].(string)
				fmt.Println("Loading data for ", fieldsName)
				dataColl := ReadLines(reCount, 61.00, fmt.Sprintf("./%s/%s.txt", dictPath, dict.(string)))
				// Rand data
				workers.ShuffleStrings(dataColl)
				// Add to dict
				dictData[fieldsName] = dataColl

			}
		}
	}
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
