package ten

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	first := First(splittedStrings)
	if first != 26397 {
		t.Errorf("First answer should be 26397, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	second := Second(splittedStrings)
	if second != 288957 {
		t.Errorf("Second answer should be 288957, result was %d", second)
	}
}
