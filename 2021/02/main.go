package main

import (
	"aoc2021"
	"fmt"
	"strconv"
	"strings"
)

func first(splitted_strings []string) int {
	horizontal, depth := 0, 0

	for _, s := range splitted_strings {
		command := strings.Split(s, " ")
		direction := command[0]
		value, err := strconv.Atoi(command[1])
		aoc2021.Check(err)
		switch direction {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		default:
			panic("Command not known")
		}
	}
	return horizontal * depth
}

func second(splitted_strings []string) int {
	horizontal, depth, aim := 0, 0, 0

	for _, s := range splitted_strings {
		command := strings.Split(s, " ")
		direction := command[0]
		value, err := strconv.Atoi(command[1])
		aoc2021.Check(err)
		switch direction {
		case "forward":
			horizontal += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			panic("Command not known")
		}
	}
	return horizontal * depth
}

func main() {
	splitted_strings := aoc2021.Read("real.txt")
	fmt.Println(first(splitted_strings))
	fmt.Println(second(splitted_strings))
}
