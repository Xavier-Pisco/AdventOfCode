package two

import (
	"2021/Utilities"
	"fmt"
	"strconv"
	"strings"
)

func First(splittedStrings []string) int {
	horizontal, depth := 0, 0

	for _, s := range splittedStrings {
		command := strings.Split(s, " ")
		direction := command[0]
		value, err := strconv.Atoi(command[1])
		Utilities.Check(err)
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

func Second(splittedStrings []string) int {
	horizontal, depth, aim := 0, 0, 0

	for _, s := range splittedStrings {
		command := strings.Split(s, " ")
		direction := command[0]
		value, err := strconv.Atoi(command[1])
		Utilities.Check(err)
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

func Solve() {
	splittedStrings := Utilities.Read("02/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
