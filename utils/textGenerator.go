package utils

import "strconv"

func GenerateTextForSolution(day int, firstProblem bool) string {
	problem := "First"
	if !firstProblem {
		problem = "Second"
	}
	return "Solution for " + problem + " Problem of Day " + strconv.Itoa(day) +" is:"
}