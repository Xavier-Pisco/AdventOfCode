package main

import (
	"aoc2021"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type line struct {
	x1, y1, x2, y2 int
}

func Max(a, b, c int) int {
	return int(math.Max(math.Max(float64(a), float64(b)), float64(c)))
}

func createDiagram(lines []line) [][]int {
	max_x, max_y := 0, 0
	for _, l := range lines {
		max_x = Max(max_x, l.x1, l.x2)
		max_y = Max(max_y, l.y1, l.y2)
	}
	diagram := make([][]int, 0)
	for i := 0; i <= max_y; i++ {
		entry := make([]int, max_x+1)
		diagram = append(diagram, entry)
	}
	return diagram
}

func getPoint(point string) (int, int) {
	p1 := strings.Split(point, ",")
	if len(p1) != 2 {
		panic("Invalid point " + point)
	}
	x1, err := strconv.Atoi(p1[0])
	aoc2021.Check(err)
	y1, err := strconv.Atoi(p1[1])
	aoc2021.Check(err)
	return x1, y1
}

func readLines(splitted_strings []string) []line {
	lines := make([]line, len(splitted_strings))
	for i, s := range splitted_strings {
		points := strings.Split(s, " -> ")
		if len(points) != 2 {
			panic("Invalid line " + s)
		}
		x1, y1 := getPoint(points[0])
		x2, y2 := getPoint(points[1])
		lines[i] = line{x1, y1, x2, y2}
	}
	return lines
}

func populateDiagramNoDiagonals(diagram *[][]int, lines []line) {
	for _, l := range lines {
		if l.x1 == l.x2 {
			min := int(math.Min(float64(l.y1), float64(l.y2)))
			max := int(math.Max(float64(l.y1), float64(l.y2)))
			for i := min; i <= max; i++ {
				(*diagram)[i][l.x1]++
			}
		} else if l.y1 == l.y2 {
			min := int(math.Min(float64(l.x1), float64(l.x2)))
			max := int(math.Max(float64(l.x1), float64(l.x2)))
			for i := min; i <= max; i++ {
				(*diagram)[l.y1][i]++
			}
		}
	}
}

func countHigher(diagram [][]int, value int) int {
	count := 0
	for _, row := range diagram {
		for _, number := range row {
			if number >= value {
				count++
			}
		}
	}
	return count
}

func first(splitted_strings []string) int {
	lines := readLines(splitted_strings)
	diagram := createDiagram(lines)
	populateDiagramNoDiagonals(&diagram, lines)
	return countHigher(diagram, 2)
}

func populateDiagram(diagram *[][]int, lines []line) {
	for _, l := range lines {
		if l.x1 == l.x2 {
			min := int(math.Min(float64(l.y1), float64(l.y2)))
			max := int(math.Max(float64(l.y1), float64(l.y2)))
			for i := min; i <= max; i++ {
				(*diagram)[i][l.x1]++
			}
		} else if l.y1 == l.y2 {
			min := int(math.Min(float64(l.x1), float64(l.x2)))
			max := int(math.Max(float64(l.x1), float64(l.x2)))
			for i := min; i <= max; i++ {
				(*diagram)[l.y1][i]++
			}
		} else {
			x_increment := int((l.x2 - l.x1) / int(math.Abs(float64(l.y2-l.y1))))
			y_increment := int((l.y2 - l.y1) / int(math.Abs(float64(l.x2-l.x1))))
			for i, j := l.x1, l.y1; i != l.x2+x_increment && j != l.y2+y_increment; i, j = i+x_increment, j+y_increment {
				(*diagram)[j][i]++
			}
		}
	}
}

func second(splitted_strings []string) int {
	lines := readLines(splitted_strings)
	diagram := createDiagram(lines)
	populateDiagram(&diagram, lines)
	return countHigher(diagram, 2)
}

func main() {
	splitted_strings := aoc2021.Read("real.txt")
	fmt.Println(first(splitted_strings))
	fmt.Println(second(splitted_strings))
}
