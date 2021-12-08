package seven

import (
	"2021/Utilities"
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

func initialPositions(splittedStrings []string) []int {
	positions := make([]int, 0)
	for _, s := range splittedStrings {
		position, err := strconv.Atoi(s)
		Utilities.Check(err)
		positions = append(positions, position)
	}
	return positions
}

func moveToPosition(positions []int, finalPosition int) int {
	totalFuel := 0
	for _, position := range positions {
		totalFuel += fuelNeeded(position, finalPosition)
	}
	return totalFuel
}

func moveToPosition2(positions []int, finalPosition int) int {
	totalFuel := 0
	for _, position := range positions {
		totalFuel += fuelNeeded2(position, finalPosition)
	}
	return totalFuel
}

func maxPosition(positions []int) int {
	max := 0
	for _, position := range positions {
		max = int(math.Max(float64(max), float64(position)))
	}
	return max
}

func First(splittedStrings []string) int {
	positions := initialPositions(splittedStrings)
	limit := maxPosition(positions)
	best := math.MaxInt
	for i := 0; i < limit; i++ {
		best = int(math.Min(float64(best), float64(moveToPosition(positions, i))))
	}
	return best
}

func Second(splittedStrings []string) int {
	positions := initialPositions(splittedStrings)
	limit := maxPosition(positions)
	best := math.MaxInt
	for i := 0; i < limit; i++ {
		best = int(math.Min(float64(best), float64(moveToPosition2(positions, i))))
	}
	return best
}

func Solve() {
	splittedStrings := Utilities.ReadWithDelimiter("07/real.txt", ",")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
