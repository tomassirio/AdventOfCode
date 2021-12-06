package day6

import (
	"strconv"
	"strings"
)

const DAYS1 = 80
const DAYS2 = 256
const REPRODUCTION_CYCLE = 9

type Day struct {}

func NewDay() Day{
	return Day{}
}

func (d Day) Solution1(input []string) int {
	return createSolution(input, DAYS1)
}

func (d Day) Solution2(input []string) int {
	return createSolution(input, DAYS2)
}

func createSolution(input []string, days int) int {
	school := fishParser(input)
	rd, tf := createReproductiveDates(school)
	var today int

	for i := 0; i < days; i ++ {
		today = i % REPRODUCTION_CYCLE

		newFish := rd[today]
		tf += newFish

		newFishReproductionDate := (today + 9) % REPRODUCTION_CYCLE
		oldFishReproductionDate := (today + 7) % REPRODUCTION_CYCLE
		rd[newFishReproductionDate] += newFish
		rd[oldFishReproductionDate] += newFish
		rd[today] -= newFish
	}
	return tf
}

func fishParser(input []string) []int{
	schoolStrings := strings.Split(input[0], ",")
	var school []int

	for _, s := range schoolStrings {
		fishInt, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		school = append(school, fishInt)
	}
	return school
}

func createReproductiveDates(school []int) ([]int, int){
	totalFish := 0
	reproductionDates := make([]int, REPRODUCTION_CYCLE)
	for _, fish := range school {
		reproductionDates[fish] ++
		totalFish ++
	}
	return reproductionDates, totalFish
}
