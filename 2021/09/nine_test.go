package nine

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	first := First(splittedStrings)
	if first != 15 {
		t.Errorf("First answer should be 15, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	second := Second(splittedStrings)
	if second != 1134 {
		t.Errorf("Second answer should be 1134, result was %d", second)
	}
}
