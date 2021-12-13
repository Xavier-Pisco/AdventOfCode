package nine

import (
	"2021/Utilities"
	"fmt"
	"math"
)

type point struct {
	i, j int
}

func smallestBasin(basins []int) (int, int) {
	position, size := 0, math.MaxInt
	for i := range basins {
		if basins[i] < size {
			position = i
			size = basins[i]
		}
	}
	return position, size
}

func calculateBasinHelper(p point, heightMap [][]int, positions *map[point]bool) {
	(*positions)[p] = true
	value := heightMap[p.i][p.j]
	if value == 8 {
		return
	}
	top, bottom, left, right := calculateNeighbours(heightMap, p.i, p.j)
	if top != math.MaxInt && top != 9 && value < top {
		calculateBasinHelper(point{p.i - 1, p.j}, heightMap, positions)
	}
	if bottom != math.MaxInt && bottom != 9 && value < bottom {
		calculateBasinHelper(point{p.i + 1, p.j}, heightMap, positions)
	}
	if left != math.MaxInt && left != 9 && value < left {
		calculateBasinHelper(point{p.i, p.j - 1}, heightMap, positions)
	}
	if right != math.MaxInt && right != 9 && value < right {
		calculateBasinHelper(point{p.i, p.j + 1}, heightMap, positions)
	}
}

func calculateBasin(p point, heightMap [][]int) int {
	positions := make(map[point]bool)
	calculateBasinHelper(p, heightMap, &positions)
	return len(positions)
}

func Second(splittedStrings []string) int {
	heightMap := createHeightMap(splittedStrings)
	lowPoints := make([]point, 0)
	for i := range heightMap {
		for j := range heightMap[i] {
			if checkLowPoint(heightMap, i, j) {
				lowPoints = append(lowPoints, point{i, j})
			}
		}
	}
	largestBasins := [3]int{0, 0, 0}
	for _, p := range lowPoints {
		basin := calculateBasin(p, heightMap)
		position, size := smallestBasin(largestBasins[:])
		if basin > size {
			largestBasins[position] = basin
		}
	}
	return largestBasins[0] * largestBasins[1] * largestBasins[2]
}

func createHeightMap(splittedStrings []string) [][]int {
	heightMap := make([][]int, 0)
	for _, line := range splittedStrings {
		row := make([]int, 0)
		for _, char := range line {
			row = append(row, Utilities.RuneToInt(char))
		}
		heightMap = append(heightMap, row)
	}

	return heightMap
}

func calculateNeighbours(heightMap [][]int, i int, j int) (int, int, int, int) {
	top, bottom, left, right := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	if i == 0 {
		bottom = heightMap[i+1][j]
		if j == 0 {
			right = heightMap[i][j+1]
		} else if j == len(heightMap[i])-1 {
			left = heightMap[i][j-1]
		} else {
			left = heightMap[i][j-1]
			right = heightMap[i][j+1]
		}
	} else if i == len(heightMap)-1 {
		top = heightMap[i-1][j]
		if j == 0 {
			right = heightMap[i][j+1]
		} else if j == len(heightMap[i])-1 {
			left = heightMap[i][j-1]
		} else {
			left = heightMap[i][j-1]
			right = heightMap[i][j+1]
		}
	} else if j == 0 {
		top = heightMap[i-1][j]
		bottom = heightMap[i+1][j]
		right = heightMap[i][j+1]
	} else if j == len(heightMap[i])-1 {
		top = heightMap[i-1][j]
		bottom = heightMap[i+1][j]
		left = heightMap[i][j-1]
	} else {
		top = heightMap[i-1][j]
		bottom = heightMap[i+1][j]
		left = heightMap[i][j-1]
		right = heightMap[i][j+1]
	}
	return top, bottom, left, right
}

func checkLowPoint(heightMap [][]int, i int, j int) bool {
	position := heightMap[i][j]
	top, bottom, left, right := calculateNeighbours(heightMap, i, j)
	return position < top && position < bottom && position < left && position < right
}

func First(splittedStrings []string) int {
	heightMap := createHeightMap(splittedStrings)
	result := 0
	for i := range heightMap {
		for j := range heightMap[i] {
			if checkLowPoint(heightMap, i, j) {
				result += heightMap[i][j] + 1
			}
		}
	}
	return result
}

func Solve() {
	splittedStrings := Utilities.Read("09/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
