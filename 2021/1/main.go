package main

import (
	"fmt"
	"math"
	"strconv"

	"aoc2021"
)

func first(splitted_strings []string) int {
	previous_value, increases := math.MaxInt, 0
	for _, s := range splitted_strings {
		value, err := strconv.Atoi(s)
		aoc2021.Check(err)
		if value > previous_value {
			increases++
		}
		previous_value = value
	}
	return increases
}

func sum_array(array []string) int {
	sum := 0
	for _, s := range array {
		value, err := strconv.Atoi(s)
		aoc2021.Check(err)
		sum += value
	}
	return sum
}

func second(splitted_strings []string) int {
	sum, increases, previous_sum := math.MaxInt, 0, math.MaxInt
	var slice []string
	for i := 0; i < len(splitted_strings)-2; i++ {
		slice = splitted_strings[i : i+3]
		previous_sum = sum
		sum = sum_array(slice)
		if previous_sum < sum {
			increases++
		}
	}
	return increases
}

func main() {
	splitted_strings := aoc2021.Read("real.txt")
	fmt.Println("First: ", first(splitted_strings))
	fmt.Println("Second: ", second(splitted_strings))
}
