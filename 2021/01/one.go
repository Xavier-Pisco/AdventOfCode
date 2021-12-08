package one

import (
	"2021/Utilities"
	"fmt"
	"math"
	"strconv"
)

func First(splittedStrings []string) int {
	previousValue, increases := math.MaxInt, 0
	for _, s := range splittedStrings {
		value, err := strconv.Atoi(s)
		Utilities.Check(err)
		if value > previousValue {
			increases++
		}
		previousValue = value
	}
	return increases
}

func sumArray(array []string) int {
	sum := 0
	for _, s := range array {
		value, err := strconv.Atoi(s)
		Utilities.Check(err)
		sum += value
	}
	return sum
}

func Second(splittedStrings []string) int {
	sum, increases, previousSum := math.MaxInt, 0, math.MaxInt
	var slice []string
	for i := 0; i < len(splittedStrings)-2; i++ {
		slice = splittedStrings[i : i+3]
		previousSum = sum
		sum = sumArray(slice)
		if previousSum < sum {
			increases++
		}
	}
	return increases
}

func Solve() {
	splittedStrings := Utilities.Read("01/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
