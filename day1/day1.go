package day1

import (
	"strconv"
)

type Day struct {}

func NewDay() Day{
	return Day{}
}

func (d Day) Solution1(input []string) int{
	previousValue := -1
	increasedValues := 0
	for _, value := range input {
		numericValue, _ := strconv.Atoi(value)
		if previousValue > 0 {
			if numericValue > previousValue {
				increasedValues++
			}
		}
		previousValue = numericValue
	}
	return increasedValues
}

func (d Day) Solution2(input []string) int{
	previousValue := -1
	increasedValues := 0
	for i := 0; i < len(input) -2; i++ {
		value1,_ := strconv.Atoi(input[i])
		value2,_ := strconv.Atoi(input[i+1])
		value3,_ := strconv.Atoi(input[i+2])
		threeValueSum := value1 + value2 + value3
		if previousValue > 0 {
			if threeValueSum > previousValue {
				increasedValues++
			}
		}
		previousValue = threeValueSum
	}
	return increasedValues
}

func (d Day) PrintSolution(){}