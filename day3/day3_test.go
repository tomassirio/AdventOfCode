package day3

import (
	"testing"
)

const SOLUTION_1_RESULT = 198
const SOLUTION_2_RESULT = 230

func TestSolution1(t *testing.T) {
	inputMock := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	functionResult := Solution1(inputMock)

	if functionResult != SOLUTION_1_RESULT {
		t.Errorf("Solution1 should return 198 under given input. Expected: %d, Gotten: %d", SOLUTION_1_RESULT, functionResult)
	}
}

func TestSolution2(t *testing.T) {
	inputMock := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	functionResult := Solution2(inputMock)

	if functionResult != SOLUTION_2_RESULT {
		t.Errorf("Solution2 should return 230 under given input. Expected: %d, Gotten: %d", SOLUTION_2_RESULT, functionResult)
	}
}