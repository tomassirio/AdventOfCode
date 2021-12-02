package main

import (
	"fmt"
	"github.com/tomassirio/adventOfCode/day1"
	"github.com/tomassirio/adventOfCode/day2"
	"github.com/tomassirio/adventOfCode/utils"
)

func main() {
	fmt.Println(utils.GenerateTextForSolution(1, true), day1.Solution1())
	fmt.Println(utils.GenerateTextForSolution(1, false), day1.Solution2())

	fmt.Println(utils.GenerateTextForSolution(2, true), day2.Solution1())
	fmt.Println(utils.GenerateTextForSolution(2, false), day2.Solution2())
}