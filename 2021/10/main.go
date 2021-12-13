package ten

import (
	"2021/Utilities"
	"fmt"
	"sort"
)

func secondPoints(char rune) int {
	if char == ')' {
		return 1
	}
	if char == ']' {
		return 2
	}
	if char == '}' {
		return 3
	}
	if char == '>' {
		return 4
	}
	panic("Unexpected rune " + string(char))
}

func Second(splittedStrings []string) int {
	incompleteLines := make([][]rune, 0)
	for _, line := range splittedStrings {
		missingChars, illegal := parseLine(line)
		if !illegal {
			incompleteLines = append(incompleteLines, missingChars)
		}
	}
	linesPoints := make([]int, 0)
	for _, line := range incompleteLines {
		points := 0
		for i := len(line) - 1; i >= 0; i-- {
			points *= 5
			points += secondPoints(line[i])
		}
		linesPoints = append(linesPoints, points)
	}
	sort.Ints(linesPoints)
	return linesPoints[len(linesPoints)/2]
}

func points(char rune) int {
	if char == ')' {
		return 3
	}
	if char == ']' {
		return 57
	}
	if char == '}' {
		return 1197
	}
	if char == '>' {
		return 25137
	}
	panic("Unexpected rune " + string(char))
}

func parseLine(line string) ([]rune, bool) {
	expectedChars := make([]rune, 0)
	for _, char := range line {
		if char == '(' {
			expectedChars = append(expectedChars, ')')
		} else if char == '[' {
			expectedChars = append(expectedChars, ']')
		} else if char == '{' {
			expectedChars = append(expectedChars, '}')
		} else if char == '<' {
			expectedChars = append(expectedChars, '>')
		} else if char == ')' || char == ']' || char == '}' || char == '>' {
			if expectedChars[len(expectedChars)-1] == char {
				expectedChars = expectedChars[:len(expectedChars)-1]
			} else {
				return []rune{char}, true
			}
		}
	}
	return expectedChars, false
}

func First(splittedStrings []string) int {
	illegalCharacters := make([]rune, 0)
	for _, line := range splittedStrings {
		char, illegal := parseLine(line)
		if illegal {
			illegalCharacters = append(illegalCharacters, char[0])
		}
	}
	result := 0
	for _, char := range illegalCharacters {
		result += points(char)
	}
	return result
}

func Solve() {
	splittedStrings := Utilities.Read("10/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
