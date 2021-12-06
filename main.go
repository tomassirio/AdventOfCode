package main

import (
	"fmt"
	"github.com/tomassirio/adventOfCode/utils"
)

const DAY_NUMBER = 6

func main() {
	day := 1
	for day <= DAY_NUMBER {
		fmt.Println(utils.PrintSolutionForDay(day))
		day++
	}
}