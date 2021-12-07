package main

import (
	"aoc2021"
	"fmt"
	"math"
	"strconv"
)

func first(splitted_strings []string) int {
	var number_bits = len(splitted_strings[0])
	var number_lines = len(splitted_strings)
	var ones_count = make([]uint, number_bits)
	for _, value := range splitted_strings {
		for i, bit := range value {
			if bit == '1' {
				ones_count[i]++
			}
		}
	}
	gamma, epsilon := 0, 0
	for i, count := range ones_count {
		if count > uint(number_lines)/2 {
			gamma += int(math.Pow(2, float64(number_bits-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(number_bits-i-1)))
		}
	}
	return gamma * epsilon
}

func most_common_bit(splitted_strings []string, position int) int {
	zeros, ones := 0, 0
	for _, value := range splitted_strings {
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

func remove_non_equals(array []string, position int, bit int) []string {
	var result []string
	for _, value := range array {
		if value[position] == strconv.Itoa(bit)[0] {
			result = append(result, value)
		}
	}
	return result
}

func until_one(splitted_strings []string, position int, keep int) []string {
	if len(splitted_strings) == 1 {
		return splitted_strings
	}
	most_common := most_common_bit(splitted_strings, position)
	if most_common == -1 {
		return until_one(remove_non_equals(splitted_strings, position, keep), position+1, keep)
	} else {
		return until_one(remove_non_equals(splitted_strings, position, (most_common+keep+1)%2), position+1, keep)
	}
}

func second(splitted_strings []string) int {
	oxygen := until_one(splitted_strings, 0, 1)
	co2 := until_one(splitted_strings, 0, 0)
	oxygen_value, co2_value := 0, 0
	for i, value := range oxygen[0] {
		oxygen_value += int(int(value-'0') * int(math.Pow(2, float64(len(oxygen[0])-i-1))))
	}
	for i, value := range co2[0] {
		co2_value += int(int(value-'0') * int(math.Pow(2, float64(len(co2[0])-i-1))))
	}
	return oxygen_value * co2_value
}

func main() {
	splitted_strings := aoc2021.Read("real.txt")
	fmt.Println(first(splitted_strings))
	fmt.Println(second(splitted_strings))
}
