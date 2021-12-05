package utils


import (
	"io/ioutil"
	"strconv"
	"strings"
)

func GetFileInput(path string) []string{
	file := OpenFile(path)
	lines := strings.Split(string(file), "\n")

	return lines
}

func OpenFile(path string) []byte {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		panic("File couldn't be read")
	}
	return content
}

func GetInput(day string) []string{
	return GetFileInput("day"+day+"/input.txt")
}

func GetCleanInput(day int) []string{
	//Remove empty space from lines
	input := GetInput(strconv.Itoa(day))
	return input[:len(input)-1]
}
