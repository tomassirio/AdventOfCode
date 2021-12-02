package utils


import (
	"io/ioutil"
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
	return GetFileInput(day+"/input.txt")
}