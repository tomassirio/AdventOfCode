package day7

import (
	"strconv"
	"strings"
)

const INT_MAX = 9999999999

type Day struct {}

func NewDay() Day{
	return Day{}
}

func (d Day) Solution1(input []string) int {
	parsedInput := parseInput(input)
	solutions := calculateDistances(parsedInput, false)
	return getResult(solutions)
}

func (d Day) Solution2(input []string) int {
	parsedInput := parseInput(input)
	solutions := calculateDistances(parsedInput, true)
	return getResult(solutions)
}

func mod(number int) int {
	if number < 0 {
		number = -number
	}
	return number
}

func parseInput(input []string) []int {
	var parsedInput []int
	str := strings.Split(input[0], ",")

	for _, value := range str {
		parsedValue, _ := strconv.Atoi(value)
		parsedInput = append(parsedInput, parsedValue)
	}
	return parsedInput
}

func fillSolutions(input []int) []int {
	solutions := make([]int, len(input))
	for i := range solutions {
		solutions[i] = 0
	}
	return solutions
}

func getResult(solutions []int) int {
	result := INT_MAX
	for _, value := range solutions {
		if value < result && value != 0{
			result = value
		}
	}
	return result
}

func calculateDistances(parsedInput []int, extraSteps bool) []int {
	solutions := fillSolutions(parsedInput)
	for i, value := range parsedInput {
		for _, position := range parsedInput {
			moddedDistance := mod(value - position)
			steps := 0
			if extraSteps {
				for j := 1; j < moddedDistance; j++ {
					steps += j
				}
			}
			solutions[i] += moddedDistance + steps
		}
	}
	return solutions
}