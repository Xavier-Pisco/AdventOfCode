package main

import (
	"aoc2021"
	"fmt"
	"strconv"
)

func initiateFishes(splitted_strings []string) map[int]int {
	fishes := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, s := range splitted_strings {
		fish, err := strconv.Atoi(s)
		aoc2021.Check(err)
		fishes[fish]++
	}
	return fishes
}

func reproduceFishes(fishes map[int]int, days int) map[int]int {
	if days == -1 {
		return fishes
	}
	new_fishes := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for i := 0; i < 8; i++ {
		new_fishes[i] = fishes[i+1]
	}
	new_fishes[8] = fishes[0]
	new_fishes[6] += fishes[0]
	return reproduceFishes(new_fishes, days-1)
}

func first(splitted_strings []string) int {
	fishes := initiateFishes(splitted_strings)
	fishes = reproduceFishes(fishes, 80)
	total := 0
	for i := 0; i < 8; i++ {
		total += fishes[i]
	}
	return total
}

func second(splitted_strings []string) int {
	fishes := initiateFishes(splitted_strings)
	fishes = reproduceFishes(fishes, 256)
	total := 0
	for i := 0; i < 8; i++ {
		total += fishes[i]
	}
	return total
}

func main() {
	splitted_strings := aoc2021.ReadWithDelimiter("real.txt", ",")
	fmt.Println(first(splitted_strings))
	fmt.Println(second(splitted_strings))
}
