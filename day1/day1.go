package day1

import (
	"github.com/tomassirio/adventOfCode/utils"
	"strconv"
)

func getInput() []string{
	return utils.ScanFile("day1/input.txt")
}

func Solution1() int{
	input := getInput()
	previousValue := -1
	increasedValues := 0
	for _, value := range input {
		//fmt.Println(value)
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

func Solution2() int{
	input := getInput()
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