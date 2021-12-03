package main

import (
	"fmt"
	"github.com/tomassirio/adventOfCode/day1"
	"github.com/tomassirio/adventOfCode/day2"
	"github.com/tomassirio/adventOfCode/day3"
	"github.com/tomassirio/adventOfCode/utils"
)

func main() {
	fmt.Println(utils.GenerateTextForSolution(1, true), day1.Solution1(utils.GetCleanInput("day1")))
	fmt.Println(utils.GenerateTextForSolution(1, false), day1.Solution2(utils.GetCleanInput("day1")))

	fmt.Println(utils.GenerateTextForSolution(2, true), day2.Solution1(utils.GetCleanInput("day2")))
	fmt.Println(utils.GenerateTextForSolution(2, false), day2.Solution2(utils.GetCleanInput("day2")))

	fmt.Println(utils.GenerateTextForSolution(3, true), day3.Solution1(utils.GetCleanInput("day3")))
	fmt.Println(utils.GenerateTextForSolution(3, false), day3.Solution2(utils.GetCleanInput("day3")))
}