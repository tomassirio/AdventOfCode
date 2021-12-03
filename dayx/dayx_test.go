package dayx

import (
	"testing"
)

const SOLUTION_1_RESULT = -1
const SOLUTION_2_RESULT = -1

func TestSolution1(t *testing.T) {
	inputMock := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	functionResult := NewDay().Solution1(inputMock)

	if functionResult != SOLUTION_1_RESULT {
		t.Errorf("Solution1 should return %d under given input. Expected: %d, Gotten: %d", SOLUTION_1_RESULT, SOLUTION_1_RESULT, functionResult)
	}
}

func TestSolution2(t *testing.T) {
	inputMock := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	functionResult := NewDay().Solution2(inputMock)

	if functionResult != SOLUTION_2_RESULT {
		t.Errorf("Solution2 should return %d under given input. Expected: %d, Gotten: %d", SOLUTION_2_RESULT, SOLUTION_2_RESULT, functionResult)
	}
}