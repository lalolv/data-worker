package handler

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// ReadLines Read a few line
func ReadLines(lines []int, path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Read file err: ", err.Error())
		return []string{}
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Read()
	// return list
	var list []string
	var line, index int
	for {
		lineText, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read file's content err: ", err.Error())
			}
			break
		}
		if lines[index] == line {
			list = append(list, lineText...)
			index++
			if index >= len(lines) {
				break
			}
		}
		line++
	}

	return list
}
