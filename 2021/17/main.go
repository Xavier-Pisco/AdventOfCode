package seventeen

import (
	"2021/Utilities"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Second(splittedStrings []string) int {
	xRange, yRange := readRanges(splittedStrings[0])
	validX := calculateX(xRange)
	validY := calculateY(yRange)
	validJumps := calculateValidJumps(validX, validY)
	return len(validJumps)
}

func readRanges(s string) ([2]int, [2]int) {
	values := strings.Split(s, "=")
	xRangeString := strings.Split(strings.Split(values[1], ",")[0], "..")
	yRangeString := strings.Split(strings.Split(values[2], " ")[0], "..")
	xRange, yRange := [2]int{}, [2]int{}
	for i := 0; i < 2; i++ {
		var err error
		xRange[i], err = strconv.Atoi(xRangeString[i])
		Utilities.Check(err)
		yRange[i], err = strconv.Atoi(yRangeString[i])
		Utilities.Check(err)
	}
	return xRange, yRange
}

func calculateX(xRange [2]int) map[int][2]int {
	validX := make(map[int][2]int)
	for i := 0; i <= xRange[1]; i++ {
		position, velocity, minValid, maxValid, step := 0, i, math.MaxInt, 0, 0
		for position <= xRange[1] {
			if position >= xRange[0] && step < minValid {
				minValid = step
			}
			if velocity == 0 {
				maxValid = math.MaxInt
				break
			}
			maxValid = step
			position += velocity
			velocity -= 1
			step++
		}
		if minValid != math.MaxInt {
			validX[i] = [2]int{minValid, maxValid}
		}
	}
	return validX
}

func calculateY(yRange [2]int) map[int][2]int {
	validX := make(map[int][2]int)
	for i := yRange[0]; i <= int(math.Abs(float64(yRange[0]))); i++ {
		position, velocity, minValid, maxValid, step := 0, i, math.MaxInt, 0, 0
		for {
			if position <= yRange[1] && position >= yRange[0] && step < minValid {
				minValid = step
			}
			if position < yRange[0] {
				maxValid = step - 1
				break
			}
			maxValid = step
			position += velocity
			velocity -= 1
			step++
		}
		if minValid != math.MaxInt {
			validX[i] = [2]int{minValid, maxValid}
		}
	}
	return validX
}

func calculateValidJumps(validX map[int][2]int, validY map[int][2]int) [][2]int {
	validJumps := make([][2]int, 0)
	for kx, vx := range validX {
		for ky, vy := range validY {
			if (vy[0] >= vx[0] && vy[0] <= vx[1]) || (vy[1] >= vx[0] && vy[1] <= vx[1]) {
				validJumps = append(validJumps, [2]int{kx, ky})
			}
		}
	}
	return validJumps
}

func sumUntilZero(num int) int {
	if num == 0 {
		return 0
	}
	return num + sumUntilZero(num-1)
}

func highestAltitude(validJumps [][2]int) int {
	highestJump := [2]int{0, math.MinInt}
	for _, jump := range validJumps {
		if jump[1] > highestJump[1] {
			highestJump = jump
		}
	}
	return sumUntilZero(highestJump[1])
}

func First(splittedStrings []string) int {
	xRange, yRange := readRanges(splittedStrings[0])
	validX := calculateX(xRange)
	validY := calculateY(yRange)
	validJumps := calculateValidJumps(validX, validY)
	return highestAltitude(validJumps)
}

func Solve() {
	splittedStrings := Utilities.Read("17/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
