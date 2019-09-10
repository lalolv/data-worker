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
	"github.com/lalolv/data-worker/utils"
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
		field, ok := v.(map[string]interface{})
		if ok {
			// read data
			fieldName := field["name"].(string)
			fieldValue := field["value"].(string)
			fmt.Println("Loading data for", fieldName)
			loadFieldData(fieldName, fieldValue, dictPath, dictMaps, reCount)
		}
	}
}

// loadFieldData 解析字典字段
// @value 每个字段的值
// @dictMaps 字典列表，用来判断类型是否是字典
func loadFieldData(fieldName, fieldValue, dictPath string, dictMaps map[string]interface{}, reCount float64) {
	i := 0
	for {
		start := IndexStart(fieldValue, "{", i)
		if start < 0 {
			break
		}
		end := IndexStart(fieldValue, "}", start)

		// 字段名称
		dictKey := fieldValue[start+1 : end]
		// 判断key是否字典
		if goutil.InStringArray(DictKeys, dictKey, nil) {
			// Get dict info from list
			dictMap := dictMaps[dictKey].(map[string]interface{})
			// Read data
			count, _ := goutil.ToFloat64(dictMap["count"])
			dataColl := ReadLines(reCount, count, fmt.Sprintf("./%s/%s", dictPath, dictMap["file"]))
			// Rand data
			workers.ShuffleStrings(dataColl)
			// Add to dict
			// key 的格式：字段名.字典名
			fullKey := fmt.Sprintf("%s.%s", fieldName, dictKey)
			dictData[fullKey] = dataColl
		}

		// 确定下次查找位置
		i = end
	}

}

// IndexStart returns the index of the first instance of sep in s from a start position,
// or -1 if sep is not present in s.
func IndexStart(s, sep string, start int) int {
	if start < 0 {
		start = 0
	}
	n := len(sep)
	size := len(s)
	if n == 0 {
		return 0
	}
	c := sep[0]
	if n == 1 {
		// special case worth making fast
		for i := start; i < size; i++ {
			if s[i] == c {
				return i
			}
		}
		return -1
	}
	// n > 1
	for i := start; i+n <= size; i++ {
		if s[i] == c && s[i:i+n] == sep {
			return i
		}
	}
	return -1
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
// reCount return count
// rowCount dict file row count
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
		time.Sleep(time.Millisecond * SleepDelay)
		p.Advance(uint(len(allText)))
	}
	// Add other text
	restIndex := int(reCount - rowCount*loopN)
	list = append(list, allText[0:restIndex]...)
	// progress
	time.Sleep(time.Millisecond * SleepDelay)
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
	utils.Seed(0)
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
			utils.Seed(0)
			index = workers.Number(b0, b0+bn)
		}

		line++
		time.Sleep(time.Millisecond * SleepDelay)
		p.Advance()
	}

	return list
}
