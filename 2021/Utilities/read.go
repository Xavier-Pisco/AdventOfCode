package Utilities

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read(path string) []string {
	return ReadWithDelimiter(path, "\n")
}

func ReadWithDelimiter(path string, delimiter string) []string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var text string = string(bytes)

	return strings.Split(text, delimiter)
}

func CompareArrays(line1 []int, line2 []int) bool {
	if len(line1) != len(line2) {
		return false
	}
	for i := range line1 {
		if line1[i] != line2[i] {
			return false
		}
	}
	return true
}

func RuneToInt(char rune) int {
	return int(char - '0')
}
