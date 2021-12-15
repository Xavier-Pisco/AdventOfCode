package fifteen

import (
	"2021/Utilities"
	"fmt"
	"math"
)

type position struct {
	x, y int
}

func replicateLeft(tile [][]int, times int, matrix *[][]int) {
	for x, line := range tile {
		finalLine := make([]int, len(line)*5)
		for i, num := range line {
			for j := 0; j < times; j++ {
				if num+j < 10 {
					finalLine[i+len(line)*j] = (num + j)
				} else {
					finalLine[i+len(line)*j] = (num+j)%10 + 1
				}

			}
		}
		(*matrix)[x] = finalLine
	}
}

func replicateBottom(tile [][]int, times int, matrix *[][]int) {
	for i := 0; i < len(tile); i++ {
		line := (*matrix)[i]
		for j := 0; j < 5; j++ {
			(*matrix)[i+j*len(tile)] = make([]int, len(line))
			for x, num := range line {
				if num+j < 10 {
					(*matrix)[i+j*len(tile)][x] = num + j
				} else {
					(*matrix)[i+j*len(tile)][x] = (num+j)%10 + 1
				}
			}
		}
	}
}

func matrixFromTile(tile [][]int) [][]int {
	matrix := make([][]int, len(tile)*5)
	replicateLeft(tile, 5, &matrix)
	replicateBottom(tile, 5, &matrix)
	return matrix
}

func Second(splittedStrings []string) int {
	tile := initializeMatrix(splittedStrings)
	matrix := matrixFromTile(tile)
	return dijkstra(matrix, position{0, 0}, position{len(matrix) - 1, len(matrix[0]) - 1})
}

func initializeMatrix(splittedStrings []string) [][]int {
	matrix := make([][]int, 0)
	for _, line := range splittedStrings {
		row := make([]int, 0)
		for _, c := range line {
			row = append(row, Utilities.RuneToInt(c))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func contais(list []position, p position) bool {
	for _, p2 := range list {
		if p == p2 {
			return true
		}
	}
	return false
}

func findMin(queue []position, dist [][]int) int {
	min, position := math.MaxInt, 0
	for i, p := range queue {
		if dist[p.y][p.x] < min {
			min = dist[p.y][p.x]
			position = i
		}
	}
	return position
}

func neighbours(p position, matrix [][]int) []position {
	neighbours := make([]position, 0)
	if p.x > 0 {
		neighbours = append(neighbours, position{p.x - 1, p.y})
	}
	if p.y > 0 {
		neighbours = append(neighbours, position{p.x, p.y - 1})
	}
	if p.x < len(matrix[p.y])-1 {
		neighbours = append(neighbours, position{p.x + 1, p.y})
	}
	if p.y < len(matrix)-1 {
		neighbours = append(neighbours, position{p.x, p.y + 1})
	}
	return neighbours
}

func dijkstra(matrix [][]int, source position, destination position) int {
	dist, prev := make([][]int, 0), make([][]position, 0)
	queue := make([]position, 0)
	for i := range matrix {
		distRow := make([]int, 0)
		prevRow := make([]position, 0)
		for j := 0; j < len(matrix[i]); j++ {
			distRow = append(distRow, 10000000)
			prevRow = append(prevRow, *new(position))
			queue = append(queue, position{j, i})
		}
		dist = append(dist, distRow)
		prev = append(prev, prevRow)
	}
	dist[source.x][source.y] = 0
	for len(queue) > 0 {
		i := findMin(queue, dist)
		u := queue[i]
		if i < len(queue)-1 {
			queue = append(queue[:i], queue[i+1:]...)
		} else {
			queue = queue[:i]
		}

		for _, v := range neighbours(u, matrix) {
			if contais(queue, v) {
				alt := dist[u.y][u.x] + matrix[v.y][v.x]
				if alt < dist[v.y][v.x] {
					dist[v.y][v.x] = alt
					prev[v.y][v.x] = u
				}
			}
		}
	}
	return dist[destination.x][destination.y]
}

func First(splittedStrings []string) int {
	matrix := initializeMatrix(splittedStrings)
	return dijkstra(matrix, position{0, 0}, position{len(matrix) - 1, len(matrix[0]) - 1})
}

func Solve() {
	splittedStrings := Utilities.Read("15/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
