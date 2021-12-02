package aoc2021

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
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var text string = string(bytes)

	return strings.Split(text, "\n")
}
