package main

import (
	"aoc2021"
	"fmt"
	"math"
	"strconv"
)

func fuelNeeded(position int, destination int) int {
	return int(math.Abs(float64(destination - position)))
}

func fuelNeeded2(position int, destination int) int {
	fuel := 0
	for i := 1; i <= int(math.Abs(float64(position-destination))); i++ {
		fuel += i
	}
	return fuel
}

func initialPositions(splitted_strings []string) []int {
	positions := make([]int, 0)
	for _, s := range splitted_strings {
		position, err := strconv.Atoi(s)
		aoc2021.Check(err)
		positions = append(positions, position)
	}
	return positions
}

func moveToPosition(positions []int, final_position int) int {
	total_fuel := 0
	for _, position := range positions {
		total_fuel += fuelNeeded(position, final_position)
	}
	return total_fuel
}

func moveToPosition2(positions []int, final_position int) int {
	total_fuel := 0
	for _, position := range positions {
		total_fuel += fuelNeeded2(position, final_position)
	}
	return total_fuel
}

func maxPosition(positions []int) int {
	max := 0
	for _, position := range positions {
		max = int(math.Max(float64(max), float64(position)))
	}
	return max
}

func first(splitted_strings []string) int {
	positions := initialPositions(splitted_strings)
	limit := maxPosition(positions)
	best := math.MaxInt
	for i := 0; i < limit; i++ {
		best = int(math.Min(float64(best), float64(moveToPosition(positions, i))))
	}
	return best
}

func second(splitted_strings []string) int {
	positions := initialPositions(splitted_strings)
	limit := maxPosition(positions)
	best := math.MaxInt
	for i := 0; i < limit; i++ {
		best = int(math.Min(float64(best), float64(moveToPosition2(positions, i))))
	}
	return best
}

func main() {
	splitted_strings := aoc2021.ReadWithDelimiter("real.txt", ",")
	fmt.Println(first(splitted_strings))
	fmt.Println(second(splitted_strings))
}
