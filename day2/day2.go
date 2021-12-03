package day2

import (
	"strconv"
	"strings"
)

type direction string

type directionPair struct {
	direction direction
	depth int
}

func newDirectionPair(direction direction, depth int) directionPair{
	return directionPair{direction, depth}
}

type position struct {
	vertical int
	horizontal int
	aim int
}

func newPosition() position{
	return position{0,0,0}
}

func generalSolution(input []string) []directionPair{
	directionPairs := createDirectionPairs(input)
	return directionPairs
}

func Solution1(input []string) int{
	return getFinalResult(calculatePosition(generalSolution(input)))
}

func Solution2(input []string) int{
	return getFinalResult(calculatePosition2(generalSolution(input)))
}

func splitFields(line string) []string{
	return strings.Split(line, " ")
}

func getDepth(fields []string) int {
	depth, err := strconv.Atoi(fields[1])
	if err != nil {
		panic("Cannot get Depth from field")
	}
	return depth
}

func createDirectionPairs(lines []string) []directionPair{
	var directionPairs []directionPair
	for _, line := range lines {
		fields := splitFields(line)
		newPair := newDirectionPair(direction(fields[0]), getDepth(fields))
		directionPairs = append(directionPairs, newPair)
	}
	return directionPairs
}

func calculatePosition(directionPairs []directionPair) position {
	position := newPosition()
	for _, directionPair := range directionPairs {
		switch directionPair.direction {
		case "forward":
			position.horizontal += directionPair.depth

		case "up":
			position.vertical -= directionPair.depth

		case "down":
			position.vertical += directionPair.depth
		}
	}
	return position
}

func calculatePosition2(directionPairs []directionPair) position {
	position := newPosition()
	for _, directionPair := range directionPairs {
		switch directionPair.direction {
		case "forward":
			position.horizontal += directionPair.depth
			position.vertical += position.aim * directionPair.depth

		case "up":
			position.aim -= directionPair.depth

		case "down":
			position.aim += directionPair.depth
		}
	}
	return position
}

func getFinalResult(position position) int {
	return position.vertical * position.horizontal
}