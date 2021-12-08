package three

import (
	"2021/Utilities"
	"fmt"
	"math"
	"strconv"
)

func First(splittedStrings []string) int {
	var numberBits = len(splittedStrings[0])
	var numberLines = len(splittedStrings)
	var onesCount = make([]uint, numberBits)
	for _, value := range splittedStrings {
		for i, bit := range value {
			if bit == '1' {
				onesCount[i]++
			}
		}
	}
	gamma, epsilon := 0, 0
	for i, count := range onesCount {
		if count > uint(numberLines)/2 {
			gamma += int(math.Pow(2, float64(numberBits-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(numberBits-i-1)))
		}
	}
	return gamma * epsilon
}

func mostCommonBit(splittedStrings []string, position int) int {
	zeros, ones := 0, 0
	for _, value := range splittedStrings {
		if value[position] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if zeros == ones {
		return -1
	} else if zeros > ones {
		return 0
	} else {
		return 1
	}
}

func removeNonEquals(array []string, position int, bit int) []string {
	var result []string
	for _, value := range array {
		if value[position] == strconv.Itoa(bit)[0] {
			result = append(result, value)
		}
	}
	return result
}

func untilOne(splittedStrings []string, position int, keep int) []string {
	if len(splittedStrings) == 1 {
		return splittedStrings
	}
	mostCommon := mostCommonBit(splittedStrings, position)
	if mostCommon == -1 {
		return untilOne(removeNonEquals(splittedStrings, position, keep), position+1, keep)
	} else {
		return untilOne(removeNonEquals(splittedStrings, position, (mostCommon+keep+1)%2), position+1, keep)
	}
}

func Second(splittedStrings []string) int {
	oxygen := untilOne(splittedStrings, 0, 1)
	co2 := untilOne(splittedStrings, 0, 0)
	oxygenValue, co2Value := 0, 0
	for i, value := range oxygen[0] {
		oxygenValue += int(int(value-'0') * int(math.Pow(2, float64(len(oxygen[0])-i-1))))
	}
	for i, value := range co2[0] {
		co2Value += int(int(value-'0') * int(math.Pow(2, float64(len(co2[0])-i-1))))
	}
	return oxygenValue * co2Value
}

func Solve() {
	splittedStrings := Utilities.Read("03/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
