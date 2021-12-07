package main

import (
	"fmt"
	"github.com/tomassirio/adventOfCode/utils"
)

func main() {
	day := 1
	for day <= utils.DAY_NUMBER {
		fmt.Println(utils.PrintSolutionForDay(day))
		day++
	}
}