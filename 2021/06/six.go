package six

import (
	"2021/Utilities"
	"fmt"
	"strconv"
)

func initiateFishes(splittedStrings []string) []int {
	fishes := []int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, s := range splittedStrings {
		fish, err := strconv.Atoi(s)
		Utilities.Check(err)
		fishes[fish]++
	}
	return fishes
}

func reproduceFishes(fishes []int, days int) []int {
	if days == -1 {
		return fishes
	}
	newFishes := []int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for i := 0; i < 8; i++ {
		newFishes[i] = fishes[i+1]
	}
	newFishes[8] = fishes[0]
	newFishes[6] += fishes[0]
	return reproduceFishes(newFishes, days-1)
}

func First(splittedStrings []string) int {
	fishes := initiateFishes(splittedStrings)
	fishes = reproduceFishes(fishes, 80)
	total := 0
	for i := 0; i < 8; i++ {
		total += fishes[i]
	}
	return total
}

func Second(splittedStrings []string) int {
	fishes := initiateFishes(splittedStrings)
	fishes = reproduceFishes(fishes, 256)
	total := 0
	for i := 0; i < 8; i++ {
		total += fishes[i]
	}
	return total
}

func Solve() {
	splittedStrings := Utilities.ReadWithDelimiter("06/real.txt", ",")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
