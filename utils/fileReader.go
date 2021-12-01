package utils


import (
	"io/ioutil"
	"strings"
)

func ScanFile(path string) []string{
	content, err := ioutil.ReadFile(path)

	if err != nil {
		panic("File couldn't be read")
	}
	lines := strings.Split(string(content), "\n")

	return lines
}
