package utils

import (
	"github.com/tomassirio/adventOfCode/day1"
	"github.com/tomassirio/adventOfCode/day2"
	"github.com/tomassirio/adventOfCode/day3"
	"github.com/tomassirio/adventOfCode/day4"
)

type iDay interface {
	Solution1([]string)int
	Solution2([]string)int
}

func DayGetter(day int) iDay {
	switch day {
	case 1: return day1.NewDay()
	case 2: return day2.NewDay()
	case 3: return day3.NewDay()
	case 4: return day4.NewDay()
	//case 5: return day1.NewDay()
	//case 6: return day1.NewDay()
	//case 7: return day1.NewDay()
	//case 8: return day1.NewDay()
	//case 9: return day1.NewDay()
	//case 10: return day1.NewDay()
	//case 12: return day1.NewDay()
	//case 13: return day1.NewDay()
	//case 14: return day1.NewDay()
	//case 15: return day1.NewDay()
	//case 16: return day1.NewDay()
	//case 17: return day1.NewDay()
	//case 18: return day1.NewDay()
	//case 19: return day1.NewDay()
	//case 20: return day1.NewDay()
	//case 21: return day1.NewDay()
	//case 22: return day1.NewDay()
	//case 23: return day1.NewDay()
	//case 24: return day1.NewDay()
	//case 25: return day1.NewDay()
	}
	return nil
}