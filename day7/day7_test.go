package day7

import (
	"testing"
)

const SOLUTION_1_RESULT = 37
const SOLUTION_2_RESULT = 168

func TestSolution1(t *testing.T) {
	inputMock := []string{
		"16,1,2,0,4,2,7,1,2,14",

	}
	functionResult := NewDay().Solution1(inputMock)

	if functionResult != SOLUTION_1_RESULT {
		t.Errorf("Solution1 should return %d under given input. Expected: %d, Gotten: %d", SOLUTION_1_RESULT, SOLUTION_1_RESULT, functionResult)
	}
}

func TestSolution2(t *testing.T) {
	inputMock := []string{
		"16,1,2,0,4,2,7,1,2,14",
	}
	functionResult := NewDay().Solution2(inputMock)

	if functionResult != SOLUTION_2_RESULT {
		t.Errorf("Solution2 should return %d under given input. Expected: %d, Gotten: %d", SOLUTION_2_RESULT, SOLUTION_2_RESULT, functionResult)
	}
}