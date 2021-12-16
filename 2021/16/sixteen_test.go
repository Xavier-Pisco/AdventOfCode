package sixteen

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	first := First(splittedStrings)
	if first != 16 {
		t.Errorf("First answer should be 16, result was %d", first)
	}
	splittedStrings = Utilities.Read("example2.txt")
	first = First(splittedStrings)
	if first != 12 {
		t.Errorf("First answer should be 12, result was %d", first)
	}
	splittedStrings = Utilities.Read("example3.txt")
	first = First(splittedStrings)
	if first != 23 {
		t.Errorf("First answer should be 23, result was %d", first)
	}
	splittedStrings = Utilities.Read("example4.txt")
	first = First(splittedStrings)
	if first != 31 {
		t.Errorf("First answer should be 31, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	second := Second(splittedStrings)
	if second != 0 {
		t.Errorf("Second answer should be 0, result was %d", second)
	}
}
