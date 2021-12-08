package four

import (
	"2021/Utilities"
	"fmt"
	"strconv"
	"strings"
)

func remove(array []string, s string) []string {
	result := make([]string, 0)
	for _, value := range array {
		if value != s {
			result = append(result, value)
		}
	}
	return result
}

func createBoards(boardStrings []string) [][][]int {
	boards := make([][][]int, 0)
	for _, board := range boardStrings {
		lines := strings.Split(board, "\n")
		convertedBoard := make([][]int, 0)
		for _, line := range lines {
			convertedLine := make([]int, 0)
			line := strings.Split(line, " ")
			line = remove(line, "")
			for _, s := range line {
				value, err := strconv.Atoi(s)
				Utilities.Check(err)
				convertedLine = append(convertedLine, value)
			}
			convertedBoard = append(convertedBoard, convertedLine)
		}
		boards = append(boards, convertedBoard)
	}
	return boards
}

func markDraw(boards *[][][]int, draw int) {
	for boardPosition, board := range *boards {
		for linePosition, line := range board {
			for position, value := range line {
				if value == draw {
					(*boards)[boardPosition][linePosition][position] = -1
				}
			}
		}
	}
}

func checkWinner(boards [][][]int) (int, bool) {
	for player, board := range boards {
		for _, line := range board {
			if Utilities.CompareArrays(line, []int{-1, -1, -1, -1, -1}) {
				return player, true
			}
		}
	}
	return -1, false
}

func findWinner(draws []string, boards [][][]int) (int, int, bool) {
	for _, draw := range draws {
		draw, err := strconv.Atoi(draw)
		Utilities.Check(err)
		markDraw(&boards, draw)
		player, winner := checkWinner(boards)
		if winner {
			return player, draw, true
		}
	}
	return -1, -1, false
}

func calculateBoardValue(board [][]int) int {
	total := 0
	for _, line := range board {
		for _, value := range line {
			if value != -1 {
				total += value
			}
		}
	}
	return total
}

func First(splittedStrings []string) int {
	draws := splittedStrings[0]
	boards := createBoards(splittedStrings[1:])
	player, draw, winner := findWinner(strings.Split(draws, ","), boards)
	if winner {
		return calculateBoardValue(boards[player]) * draw
	}
	return -1
}

func checkLoser(boards [][][]int) (int, bool) {
	winners := make(map[int]bool)
	loser := -1
	for player, board := range boards {
		for i, line := range board {
			if Utilities.CompareArrays(line, []int{-1, -1, -1, -1, -1}) {
				winners[player] = true
			}
			if Utilities.CompareArrays([]int{board[0][i], board[1][i], board[2][i], board[3][i], board[4][i]}, []int{-1, -1, -1, -1, -1}) {
				winners[player] = true
			}
		}
	}
	for player := range boards {
		if !winners[player] {
			loser = player
		}
	}
	return loser, len(winners) == len(boards)

}

func findLoser(draws []string, boards [][][]int) (int, int, bool) {
	previousLoser := 0
	for _, draw := range draws {
		draw, err := strconv.Atoi(draw)
		Utilities.Check(err)
		markDraw(&boards, draw)
		player, loser := checkLoser(boards)
		if loser {
			return previousLoser, draw, true
		}
		previousLoser = player
	}
	return -1, -1, false
}

func Second(splittedStrings []string) int {
	draws := splittedStrings[0]
	boards := createBoards(splittedStrings[1:])
	player, draw, loser := findLoser(strings.Split(draws, ","), boards)
	if loser {
		return calculateBoardValue(boards[player]) * draw
	}
	return -1

}

func Solve() {
	splittedStrings := Utilities.ReadWithDelimiter("04/real.txt", "\n\n")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
