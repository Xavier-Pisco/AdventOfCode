package twelve

import (
	"2021/Utilities"
	"fmt"
	"strings"
)

type path struct {
	begin, end string
}

func searchPathsAux2(paths []path, position string, placesVisited map[string]bool, extraVisit bool) int {
	if placesVisited[position] && Utilities.IsLower(position) {
		if position == "start" || position == "end" || extraVisit {
			return 0
		} else {
			extraVisit = true
		}
	}
	if position == "end" {
		return 1
	}
	placesVisited[position] = true
	count := 0
	for _, p := range paths {
		if p.begin == position {
			count += searchPathsAux2(paths, p.end, copy(placesVisited), extraVisit)
		} else if p.end == position {
			count += searchPathsAux2(paths, p.begin, copy(placesVisited), extraVisit)
		}
	}
	return count
}

func searchPaths2(paths []path, p path) int {
	placesVisited := make(map[string]bool)
	placesVisited["start"] = true
	count := 0
	if p.begin == "start" {
		count += searchPathsAux2(paths, p.end, copy(placesVisited), false)
	} else if p.end == "start" {
		count += searchPathsAux2(paths, p.begin, copy(placesVisited), false)
	}
	return count
}

func Second(splittedStrings []string) int {
	paths := initializePaths(splittedStrings)
	count := 0
	for _, p := range paths {
		if p.begin == "start" || p.end == "start" {
			count += searchPaths2(paths, p)
		}
	}
	return count
}

func copy(original map[string]bool) map[string]bool {
	copy := make(map[string]bool)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

func initializePaths(splittedStrings []string) []path {
	paths := make([]path, 0)
	for _, line := range splittedStrings {
		splittedLine := strings.Split(line, "-")
		paths = append(paths, path{splittedLine[0], splittedLine[1]})
	}
	return paths
}

func searchPathsAux(paths []path, position string, placesVisited map[string]bool) int {
	if placesVisited[position] && Utilities.IsLower(position) {
		return 0
	}
	if position == "end" {
		return 1
	}
	placesVisited[position] = true
	count := 0
	for _, p := range paths {
		if p.begin == position {
			count += searchPathsAux(paths, p.end, copy(placesVisited))
		} else if p.end == position {
			count += searchPathsAux(paths, p.begin, copy(placesVisited))
		}
	}
	return count
}

func searchPaths(paths []path, p path) int {
	placesVisited := make(map[string]bool)
	placesVisited["start"] = true
	count := 0
	if p.begin == "start" {
		count += searchPathsAux(paths, p.end, copy(placesVisited))
	} else if p.end == "start" {
		count += searchPathsAux(paths, p.begin, copy(placesVisited))
	}
	return count
}

func First(splittedStrings []string) int {
	paths := initializePaths(splittedStrings)
	count := 0
	for _, p := range paths {
		if p.begin == "start" || p.end == "start" {
			count += searchPaths(paths, p)
		}
	}
	return count
}

func Solve() {
	splittedStrings := Utilities.Read("12/real.txt")
	fmt.Println(First(splittedStrings))
	fmt.Println(Second(splittedStrings))
}
