package main

import (
	"data-worker/workers"
	"fmt"
)

func main() {

	workers.Seed(0)
	num := workers.Number(1, 99)
	fmt.Println(num)

}
