package thirteen

import (
	"2021/Utilities"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Second(splittedStrings []string) int {
	dotsPositions := strings.Split(splittedStrings[0], "\n")
	folds := strings.Split(splittedStrings[1], "\n")
	paper := initiatePaper(dotsPositions)
	for _, fold := range folds {
		axis, value := parseFold(fold)
		paper = foldPaper(paper, axis, value)
	}
	for _, line := range paper {
		fmt.Println(string(line))
	}
	return countVisible(paper)
}

func initiatePaper(dotsPositions []string) [][]rune {
	largestX, largestY := 0, 0
	for _, dot := range dotsPositions {
		position := strings.Split(dot, ",")
		x, err := strconv.Atoi(position[0])
		Utilities.Check(err)
		y, err := strconv.Atoi(position[1])
		Utilities.Check(err)
		largestX = int(math.Max(float64(x), float64(largestX)))
		largestY = int(math.Max(float64(y), float64(largestY)))
	}
	paper := make([][]rune, largestY+1)
	for i := range paper {
		paper[i] = make([]rune, largestX+1)
		for j := range paper[i] {
			paper[i][j] = ' '
		}
	}
	for _, dot := range dotsPositions {
		position := strings.Split(dot, ",")
		x, err := strconv.Atoi(position[0])
		Utilities.Check(err)
		y, err := strconv.Atoi(position[1])
		Utilities.Check(err)
		paper[y][x] = '#'
	}
	return paper
}

func parseFold(fold string) (rune, int) {
	fold = strings.Split(fold, "fold along ")[1]
	parts := strings.Split(fold, "=")
	axis := []rune(parts[0])[0]
	value, err := strconv.Atoi(parts[1])
	Utilities.Check(err)
	return axis, value
}

func foldHorizontal(paper [][]rune, value int) [][]rune {
	newPaper := make([][]rune, 0)
	for i := range paper {
		line := make([]rune, 0)
		for j := 0; j < value; j++ {
			if paper[i][j] == '#' || paper[i][len(paper[i])-j-1] == '#' {
				line = append(line, '#')
			} else {
				line = append(line, ' ')
			}
		}
		newPaper = append(newPaper, line)
	}
	return newPaper
}

func foldVertical(paper [][]rune, value int) [][]rune {
	newPaper := make([][]rune, 0)
	for i := 0; i < value; i++ {
		line := make([]rune, 0)
		for j := 0; j < len(paper[i]); j++ {
			if paper[i][j] == '#' || paper[len(paper)-i-1][j] == '#' {
				line = append(line, '#')
			} else {
				line = append(line, ' ')
			}
		}
		newPaper = append(newPaper, line)
	}
	return newPaper
}

func foldPaper(paper [][]rune, axis rune, value int) [][]rune {
	if axis == 'x' {
		fmt.Println(value, len(paper[0]))
		return foldHorizontal(paper, value)
	} else {
		fmt.Println(value, len(paper))
		return foldVertical(paper, value)
	}
}

func countVisible(paper [][]rune) int {
	count := 0
	for _, line := range paper {
		for _, char := range line {
			if char == '#' {
				count++
			}
		}
	}
	return count
}

func First(splittedStrings []string) int {
	dotsPositions := strings.Split(splittedStrings[0], "\n")
	folds := strings.Split(splittedStrings[1], "\n")
	paper := initiatePaper(dotsPositions)
	axis, value := parseFold(folds[0])
	paper = foldPaper(paper, axis, value)
	return countVisible(paper)
}

func Solve() {
	splittedStrings := Utilities.ReadWithDelimiter("13/real.txt", "\n\n")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
