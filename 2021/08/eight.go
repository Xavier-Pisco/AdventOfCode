package eight

import (
	"2021/Utilities"
	"fmt"
	"strings"
)

func contains(list []rune, char rune) bool {
	for _, c := range list {
		if c == char {
			return true
		}
	}
	return false
}

func addCommons(list []rune, number []rune) []rune {
	if len(list) == 0 {
		list = number
		return list
	}
	for i, c := range list {
		if !contains(number, c) {
			if i < len(list)-1 {
				list = append(list[:i], list[i+1:]...)
			} else {
				list = list[:i]
			}
		}
	}
	return list
}

func decodeNumber(number []rune, encode *map[rune][]rune, numbers [][]rune) {
	switch len(number) {
	case len(numbers[1]):
		for _, c := range numbers[1] {
			(*encode)[c] = addCommons((*encode)[c], number)
		}
	case len(numbers[4]):
		for _, c := range numbers[4] {
			(*encode)[c] = addCommons((*encode)[c], number)
		}
	case len(numbers[7]):
		for _, c := range numbers[7] {
			(*encode)[c] = addCommons((*encode)[c], number)
		}
	case len(numbers[8]):
		for _, c := range numbers[8] {
			(*encode)[c] = addCommons((*encode)[c], number)
		}
	}
}

func Second(splittedStrings []string) int {
	_, fourDigits := parseInput(splittedStrings)
	numbers := createNumbers()
	encode := make(map[rune][]rune)
	for _, fourDigit := range fourDigits {
		for _, digit := range fourDigit {
			decodeNumber([]rune(digit), &encode, numbers)
		}
	}
	fmt.Println(encode)
	return 1
}

func parseLine(line string) ([]string, []string) {
	splittedLine := strings.Split(line, " | ")
	if len(splittedLine) != 2 {
		panic("Unexpected line " + line)
	}

	uniqueSingal := strings.Split(splittedLine[0], " ")
	fourDigitOutput := strings.Split(splittedLine[1], " ")
	return uniqueSingal, fourDigitOutput
}

func parseInput(splittedStrings []string) ([][]string, [][]string) {
	totalUniqueSignals, totalFourDigit := make([][]string, 0), make([][]string, 0)
	for _, s := range splittedStrings {
		uniqueSignal, fourDigitOutput := parseLine(s)
		totalUniqueSignals = append(totalUniqueSignals, uniqueSignal)
		totalFourDigit = append(totalFourDigit, fourDigitOutput)
	}
	return totalUniqueSignals, totalFourDigit
}

func createNumbers() [][]rune {
	numbers := make([][]rune, 0)
	zero := []rune{'a', 'b', 'c', 'e', 'f', 'g'}
	numbers = append(numbers, zero)
	one := []rune{'c', 'f'}
	numbers = append(numbers, one)
	two := []rune{'a', 'c', 'd', 'e', 'g'}
	numbers = append(numbers, two)
	three := []rune{'a', 'c', 'd', 'f', 'g'}
	numbers = append(numbers, three)
	four := []rune{'b', 'c', 'd', 'f'}
	numbers = append(numbers, four)
	five := []rune{'a', 'b', 'd', 'f', 'g'}
	numbers = append(numbers, five)
	six := []rune{'a', 'b', 'd', 'e', 'f', 'g'}
	numbers = append(numbers, six)
	seven := []rune{'a', 'c', 'f'}
	numbers = append(numbers, seven)
	eight := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	numbers = append(numbers, eight)
	nine := []rune{'a', 'b', 'c', 'd', 'f', 'g'}
	numbers = append(numbers, nine)
	return numbers
}

func First(splittedStrings []string) int {
	_, fourDigits := parseInput(splittedStrings)
	numbers := createNumbers()
	count := 0
	for _, fourDigit := range fourDigits {
		for _, digit := range fourDigit {
			switch len(digit) {
			case len(numbers[1]):
				count++
			case len(numbers[4]):
				count++
			case len(numbers[7]):
				count++
			case len(numbers[8]):
				count++
			default:
				continue
			}
		}
	}
	return count
}

func Solve() {
	splittedStrings := Utilities.Read("08/example.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
