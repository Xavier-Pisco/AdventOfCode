package fourteen

import (
	"2021/Utilities"
	"fmt"
	"math"
	"strings"
)

func initializeStringPairs(template string) map[string]int {
	stringPairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		stringPairs[template[i:i+2]]++
	}
	return stringPairs
}

func secondStep(stringPairs map[string]int, pairInsertions map[string]rune) map[string]int {
	newStringPairs := make(map[string]int)
	for k, v := range stringPairs {
		newStringPairs[string(k[0])+string(pairInsertions[k])] += v
		newStringPairs[string(pairInsertions[k])+string(k[1])] += v
	}
	return newStringPairs
}

func calculateValues(stringPairs map[string]int) map[rune]int {
	count := make(map[rune]int)
	for k, v := range stringPairs {
		count[rune(k[0])] += v
		count[rune(k[1])] += v
	}
	for k, v := range count {
		count[k] = (v + 1) / 2
	}
	return count
}

func Second(splittedStrings []string) int {
	template := splittedStrings[0]
	pairInsertions := parsePairs(strings.Split(splittedStrings[1], "\n"))
	stringPairs := initializeStringPairs(template)
	for i := 0; i < 40; i++ {
		stringPairs = secondStep(stringPairs, pairInsertions)
	}
	count := calculateValues(stringPairs)
	max, min := maxAndMin(count)
	return max - min
}

func parsePairs(lines []string) map[string]rune {
	pairInsertions := make(map[string]rune)
	for _, line := range lines {
		pairParts := strings.Split(line, " -> ")
		pairInsertions[pairParts[0]] = rune(pairParts[1][0])
	}
	return pairInsertions
}

func step(template string, pairInsertions map[string]rune) string {
	newTemplate := ""
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		newTemplate += string(template[i]) + string(pairInsertions[pair])
	}
	newTemplate += string(template[len(template)-1])
	return newTemplate
}

func countLetters(template string) map[rune]int {
	count := make(map[rune]int)
	for _, c := range template {
		count[c]++
	}
	return count
}

func maxAndMin(count map[rune]int) (int, int) {
	max, min := math.MinInt, math.MaxInt
	for _, c := range count {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	return max, min
}

func First(splittedStrings []string) int {
	template := splittedStrings[0]
	pairInsertions := parsePairs(strings.Split(splittedStrings[1], "\n"))
	for i := 0; i < 10; i++ {
		template = step(template, pairInsertions)
	}
	count := countLetters(template)
	max, min := maxAndMin(count)
	return max - min
}

func Solve() {
	splittedStrings := Utilities.ReadWithDelimiter("14/real.txt", "\n\n")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
