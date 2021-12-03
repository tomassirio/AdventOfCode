package day3

import (
	"fmt"
	"github.com/tomassirio/adventOfCode/utils"
	"strconv"
)

var input = utils.GetCleanInput("day3")

func Solution1() int {
	binaryString := getBinaryWithMoreRepeatedBits()

	gammaRate := getRate(binaryString)
	epsilonRate := getRate(xor(binaryString))

	return getFinalResult(gammaRate, epsilonRate)
}

func Solution2() int {

	oxygenGenerationRating := getOxygenRate()

	c02ScrubberRating := getC02Rate()


	return getFinalResult(oxygenGenerationRating, c02ScrubberRating)
}

func xor(binary string) string {
	newSolution := ""
	for _, value := range binary {
		if value == '0' {
			newSolution += "1"
		} else {
			newSolution += "0"
		}
	}
	return newSolution
}

func getValuesLength() int {
	return len(input[0])
}

func getBinaryWithMoreRepeatedBits() string{
	valuesLength := getValuesLength()
	j := 0
	solution := ""
	for j  < valuesLength {
		solution += returnSolutionBit(getAmountOf0sAnd1s(j, input))
		j++
	}
	return solution
}

func returnSolutionBit(amountOf0s, amountOf1s int) string{
	solution := "1"
	if amountOf0s > amountOf1s {
		solution = "0"
	}
	return solution
}

func getAmountOf0sAnd1s(index int, values []string) (int, int){
	amountOf0s := 0
	amountOf1s := 0
	for _, value := range values {
		if value[index] == '0' {
			amountOf0s ++
		} else {
			amountOf1s ++
		}
	}
	return amountOf0s, amountOf1s
}

func getRate(binaryString string) int{
	rate, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		fmt.Println(err)

	}
	return int(rate)
}

func getFinalResult(rate1, rate2 int) int {
	return rate1 * rate2
}

func getOxygenRate() int{
	valuesLength := getValuesLength()
	index := 0
	newInput := copyInputValues()
	for index < valuesLength {
		temp := newInput[:0]

		amountOf0s, amountOf1s := getAmountOf0sAnd1s(index, newInput)

		if amountOf0s > amountOf1s {
			temp = pickValuesWithBit(newInput, temp, index, '0')
		} else {
			temp = pickValuesWithBit(newInput, temp, index, '1')
		}

		if len(newInput) == 1 {
			break
		}
		newInput = temp
		index++
	}
	return getRate(newInput[0])
}

func getC02Rate() int{
	valuesLength := getValuesLength()
	index := 0
	newInput := copyInputValues()
	for index < valuesLength {
		temp := newInput[:0]
		amountOf0s, amountOf1s := getAmountOf0sAnd1s(index, newInput)
		if amountOf0s <= amountOf1s {
			temp = pickValuesWithBit(newInput, temp, index, '0')
		} else {
			temp = pickValuesWithBit(newInput, temp, index, '1')
		}
		if len(newInput) == 1 {
			break
		}

		newInput = temp
		index++
	}
	return getRate(newInput[0])
}

func pickValuesWithBit(newInput, temp []string, index int, bit byte) []string{
	for _, binaryString := range newInput {
		if binaryString[index] == bit {
			temp = append(temp, binaryString)
		}
	}
	return temp
}

func copyInputValues() []string{
	newInput := make([]string, len(input))
	copy(newInput, input)
	return newInput
}
