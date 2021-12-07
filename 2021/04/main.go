package main

import (
	"aoc2021"
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

func createBoards(board_strings []string) [][][]int {
	boards := make([][][]int, 0)
	for _, board := range board_strings {
		lines := strings.Split(board, "\n")
		converted_board := make([][]int, 0)
		for _, line := range lines {
			converted_line := make([]int, 0)
			line := strings.Split(line, " ")
			line = remove(line, "")
			for _, s := range line {
				value, err := strconv.Atoi(s)
				aoc2021.Check(err)
				converted_line = append(converted_line, value)
			}
			converted_board = append(converted_board, converted_line)
		}
		boards = append(boards, converted_board)
	}
	return boards
}

func markDraw(boards *[][][]int, draw int) {
	for board_position, board := range *boards {
		for line_position, line := range board {
			for position, value := range line {
				if value == draw {
					(*boards)[board_position][line_position][position] = -1
				}
			}
		}
	}
}

func checkWinner(boards [][][]int) (int, bool) {
	for player, board := range boards {
		for _, line := range board {
			if aoc2021.CompareArrays(line, []int{-1, -1, -1, -1, -1}) {
				return player, true
			}
		}
	}
	return -1, false
}

func findWinner(draws []string, boards [][][]int) (int, int, bool) {
	for _, draw := range draws {
		draw, err := strconv.Atoi(draw)
		aoc2021.Check(err)
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

func first(splitted_strings []string) int {
	draws := splitted_strings[0]
	boards := createBoards(splitted_strings[1:])
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
			if aoc2021.CompareArrays(line, []int{-1, -1, -1, -1, -1}) {
				winners[player] = true
			}
			if aoc2021.CompareArrays([]int{board[0][i], board[1][i], board[2][i], board[3][i], board[4][i]}, []int{-1, -1, -1, -1, -1}) {
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
	previous_loser := 0
	for _, draw := range draws {
		draw, err := strconv.Atoi(draw)
		aoc2021.Check(err)
		markDraw(&boards, draw)
		player, loser := checkLoser(boards)
		if loser {
			return previous_loser, draw, true
		}
		previous_loser = player
	}
	return -1, -1, false
}

func second(splitted_strings []string) int {
	draws := splitted_strings[0]
	boards := createBoards(splitted_strings[1:])
	player, draw, loser := findLoser(strings.Split(draws, ","), boards)
	if loser {
		return calculateBoardValue(boards[player]) * draw
	}
	return -1

}

func main() {
	splitted_strings := aoc2021.ReadWithDelimiter("real.txt", "\n\n")
	fmt.Println(first(splitted_strings))
	fmt.Println(second(splitted_strings))
}
