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
