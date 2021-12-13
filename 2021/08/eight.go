package eight

import (
	"2021/Utilities"
	"fmt"
	"math"
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

func containsList(list []rune, chars []rune) bool {
	return containsNFromList(list, chars, len(chars))
}

func containsNFromList(list []rune, chars []rune, N int) bool {
	count := 0
	for _, char := range chars {
		if contains(list, char) {
			count++
		}
	}
	return count == N
}

func findByLength(numbers []string, originalNumbers [][]rune, encodedNumbers *map[int][]rune) {
	for _, number := range numbers {
		if len(number) == len(originalNumbers[1]) {
			(*encodedNumbers)[1] = []rune(number)
		} else if len(number) == len(originalNumbers[4]) {
			(*encodedNumbers)[4] = []rune(number)
		} else if len(number) == len(originalNumbers[7]) {
			(*encodedNumbers)[7] = []rune(number)
		} else if len(number) == len(originalNumbers[8]) {
			(*encodedNumbers)[8] = []rune(number)
		}
	}
}

func findOthers(numbers []string, originalNumbers [][]rune, encodedNumbers *map[int][]rune) {
	for _, number := range numbers {
		if len(number) == len(originalNumbers[2]) {
			if containsList([]rune(number), (*encodedNumbers)[1]) {
				(*encodedNumbers)[3] = []rune(number)
			} else if containsNFromList((*encodedNumbers)[4], []rune(number), 2) {
				(*encodedNumbers)[2] = []rune(number)
			} else {
				(*encodedNumbers)[5] = []rune(number)
			}
		} else if len(number) == len(originalNumbers[0]) {
			if containsList([]rune(number), (*encodedNumbers)[4]) {
				(*encodedNumbers)[9] = []rune(number)
			} else if containsList([]rune(number), (*encodedNumbers)[7]) {
				(*encodedNumbers)[0] = []rune(number)
			} else {
				(*encodedNumbers)[6] = []rune(number)
			}
		}
	}
}

func findNumbers(numbers []string, encodedNumbers *map[int][]rune, originalNumbers [][]rune) {
	findByLength(numbers, originalNumbers, encodedNumbers)
	findOthers(numbers, originalNumbers, encodedNumbers)
}

func findNumber(number string, encodedNumbers map[int][]rune) int {
	for i := 0; i < 10; i++ {
		if containsList(encodedNumbers[i], []rune(number)) && containsList([]rune(number), encodedNumbers[i]) {
			return i
		}
	}
	panic("Number not found")
}

func Second(splittedStrings []string) int {
	uniqueSignals, fourDigits := parseInput(splittedStrings)
	numbers := createNumbers()
	total := 0
	for i := 0; i < len(uniqueSignals); i++ {
		encodedNumbers := make(map[int][]rune)
		findNumbers(uniqueSignals[i], &encodedNumbers, numbers)
		value := 0
		for j, fourDigit := range fourDigits[i] {
			value += findNumber(fourDigit, encodedNumbers) * int(math.Pow(10, float64(3-j)))
		}
		total += value
	}
	return total
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
	splittedStrings := Utilities.Read("08/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
