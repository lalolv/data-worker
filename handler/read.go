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
	// Empty
	empty := map[string]interface{}{}
	// List file path
	listFile := fmt.Sprintf("./%s/.list.json", dictPath)
	// Check path exist
	if !CheckPathIsExist(listFile) {
		fmt.Println("No find dict list file!")
		return empty
	}

	data, err := ioutil.ReadFile(listFile)
	if err != nil {
		return empty
	}

	var dictMap map[string]interface{}
	err = json.Unmarshal(data, &dictMap)
	if err != nil {
		return empty
	}

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

	// progress bar
	p := progress.Bar(int(rowCount))
	p.Start()

	var list []string
	if reCount >= rowCount {
		list = readByLoop(ff, p, reCount, rowCount)
	} else {
		list = readBySection(ff, p, reCount, rowCount)
	}

	p.Finish()

	return list
}

// if reCount >= rowCount
func readByLoop(ff *os.File, p *progress.Progress, reCount, rowCount float64) []string {
	var list []string
	// scan line
	var allText []string
	scanner := bufio.NewScanner(ff)
	for scanner.Scan() {
		allText = append(allText, scanner.Text())
	}
	// Add N all text
	loopN := math.Floor(reCount / rowCount)
	for i := 0; i < int(loopN); i++ {
		list = append(list, allText...)
		// progress
		time.Sleep(time.Millisecond * SLEEP_DELAY)
		p.Advance(uint(len(allText)))
	}
	// Add other text
	restIndex := int(reCount - rowCount*loopN)
	list = append(list, allText[0:restIndex]...)
	// progress
	time.Sleep(time.Millisecond * SLEEP_DELAY)
	p.Advance(uint(restIndex))

	return list
}

// By section
// if reCount < rowCount
func readBySection(ff *os.File, p *progress.Progress, reCount, rowCount float64) []string {
	// a block count
	b0 := 0
	bn := int(math.Floor(rowCount/reCount)) - 1
	// rand index
	workers.Seed(0)
	index := workers.Number(b0, b0+bn)

	// scan line
	scanner := bufio.NewScanner(ff)

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
		time.Sleep(time.Millisecond * SLEEP_DELAY)
		p.Advance()
	}

	return list
}
