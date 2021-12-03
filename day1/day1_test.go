package day1

import (
	"testing"
)

const SOLUTION_1_RESULT = 7
const SOLUTION_2_RESULT = 5

func TestSolution1(t *testing.T) {
	inputMock := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	functionResult := Solution1(inputMock)

	if functionResult != SOLUTION_1_RESULT {
		t.Errorf("Solution1 should return %d under given input. Expected: %d, Gotten: %d", SOLUTION_1_RESULT, SOLUTION_1_RESULT, functionResult)
	}
}

func TestSolution2(t *testing.T) {
	inputMock := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	functionResult := Solution2(inputMock)

	if functionResult != SOLUTION_2_RESULT {
		t.Errorf("Solution2 should return %d under given input. Expected: %d, Gotten: %d", SOLUTION_2_RESULT, SOLUTION_2_RESULT, functionResult)
	}
}