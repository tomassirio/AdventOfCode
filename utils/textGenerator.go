package utils

import (
	"strconv"
	"strings"
)

func PrintSolutionForDay(day int) string {
	sb := strings.Builder{}

	sb.Write([]byte("Solution for The First problem of the Day "))
	sb.Write([]byte(strconv.Itoa(day)))
	sb.Write([]byte(" is: "))
	sb.Write([]byte(strconv.Itoa(DayGetter(day).Solution1(GetCleanInput(day)))))
	sb.Write([]byte("\n"))
	sb.Write([]byte("Solution for The Second problem of the Day "))
	sb.Write([]byte(strconv.Itoa(day)))
	sb.Write([]byte(" is: "))
	sb.Write([]byte(strconv.Itoa(DayGetter(day).Solution2(GetCleanInput(day)))))
	sb.Write([]byte("\n"))
	return sb.String()
}