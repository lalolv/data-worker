package main

import (
	"fmt"

	"github.com/lalolv/data-worker/workers"
)

func main() {

	workers.Seed(0)
	num := workers.Number(1, 99)
	fmt.Println(num)

}
