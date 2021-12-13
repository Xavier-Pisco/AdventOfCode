package eleven

import (
	"2021/Utilities"
	"fmt"
)

func Second(splittedStrings []string) int {
	grid := makeGrid(splittedStrings)
	i := 1
	for {
		increaseValues(&grid)
		flashed := make(map[point]bool)
		previousLen := len(flashed)
		for {
			flashGrid(&grid, &flashed)
			if previousLen == len(flashed) {
				break
			}
			previousLen = len(flashed)
		}
		clearGrid(&grid)
		if len(flashed) == len(grid)*len(grid[0]) {
			return i
		}
		i++
	}
}

func makeGrid(splittedStrings []string) [][]int {
	grid := make([][]int, 0)
	for _, line := range splittedStrings {
		row := make([]int, 0)
		for _, char := range line {
			row = append(row, Utilities.RuneToInt(char))
		}
		grid = append(grid, row)
	}
	return grid
}

func flashUp(grid *[][]int, i int, j int) {
	(*grid)[i-1][j]++
}

func flashUpLeft(grid *[][]int, i int, j int) {
	(*grid)[i-1][j-1]++
}

func flashUpRight(grid *[][]int, i int, j int) {
	(*grid)[i-1][j+1]++
}

func flashLeft(grid *[][]int, i int, j int) {
	(*grid)[i][j-1]++
}

func flashRight(grid *[][]int, i int, j int) {
	(*grid)[i][j+1]++
}

func flashDown(grid *[][]int, i int, j int) {
	(*grid)[i+1][j]++
}

func flashDownLeft(grid *[][]int, i int, j int) {
	(*grid)[i+1][j-1]++
}

func flashDownRight(grid *[][]int, i int, j int) {
	(*grid)[i+1][j+1]++
}

func flash(grid *[][]int, i int, j int) {
	if i == 0 {
		flashDown(grid, i, j)
		if j == 0 {
			flashRight(grid, i, j)
			flashDownRight(grid, i, j)
		} else if j == len((*grid)[i])-1 {
			flashDownLeft(grid, i, j)
			flashLeft(grid, i, j)
		} else {
			flashRight(grid, i, j)
			flashDownRight(grid, i, j)
			flashDownLeft(grid, i, j)
			flashLeft(grid, i, j)
		}
	} else if i == len(*grid)-1 {
		flashUp(grid, i, j)
		if j == 0 {
			flashRight(grid, i, j)
			flashUpRight(grid, i, j)
		} else if j == len((*grid)[i])-1 {
			flashLeft(grid, i, j)
			flashUpLeft(grid, i, j)
		} else {
			flashRight(grid, i, j)
			flashUpRight(grid, i, j)
			flashUpLeft(grid, i, j)
			flashLeft(grid, i, j)
		}
	} else if j == 0 {
		flashUp(grid, i, j)
		flashUpRight(grid, i, j)
		flashRight(grid, i, j)
		flashDownRight(grid, i, j)
		flashDown(grid, i, j)
	} else if j == len((*grid)[i])-1 {
		flashUp(grid, i, j)
		flashUpLeft(grid, i, j)
		flashLeft(grid, i, j)
		flashDownLeft(grid, i, j)
		flashDown(grid, i, j)
	} else {
		flashUp(grid, i, j)
		flashUpRight(grid, i, j)
		flashRight(grid, i, j)
		flashDownRight(grid, i, j)
		flashDown(grid, i, j)
		flashDownLeft(grid, i, j)
		flashLeft(grid, i, j)
		flashUpLeft(grid, i, j)
	}
}

func flashGrid(grid *[][]int, flashed *map[point]bool) {
	for i, row := range *grid {
		for j, value := range row {
			if value > 9 && !(*flashed)[point{i, j}] {
				flash(grid, i, j)
				(*flashed)[point{i, j}] = true
			}
		}
	}
}

func increaseValues(grid *[][]int) {
	for i, row := range *grid {
		for j := range row {
			(*grid)[i][j]++
		}
	}
}

type point struct {
	i, j int
}

func clearGrid(grid *[][]int) {
	for i := range *grid {
		for j := range (*grid)[i] {
			if ((*grid)[i][j]) > 9 {
				(*grid)[i][j] = 0
			}
		}
	}
}

func First(splittedStrings []string) int {
	grid := makeGrid(splittedStrings)
	count := 0
	for i := 0; i < 100; i++ {
		increaseValues(&grid)
		flashed := make(map[point]bool)
		previousLen := len(flashed)
		for {
			flashGrid(&grid, &flashed)
			if previousLen == len(flashed) {
				break
			}
			previousLen = len(flashed)
		}
		clearGrid(&grid)
		count += len(flashed)
	}
	return count
}

func Solve() {
	splittedStrings := Utilities.Read("11/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
