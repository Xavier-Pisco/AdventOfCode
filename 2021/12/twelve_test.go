package twelve

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	first := First(splittedStrings)
	if first != 10 {
		t.Errorf("First answer should be 10, result was %d", first)
	}
	splittedStrings = Utilities.Read("example2.txt")
	first = First(splittedStrings)
	if first != 19 {
		t.Errorf("First answer should be 19, result was %d", first)
	}
	splittedStrings = Utilities.Read("example3.txt")
	first = First(splittedStrings)
	if first != 226 {
		t.Errorf("First answer should be 226, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	second := Second(splittedStrings)
	if second != 36 {
		t.Errorf("Second answer should be 36, result was %d", second)
	}
	splittedStrings = Utilities.Read("example2.txt")
	second = Second(splittedStrings)
	if second != 103 {
		t.Errorf("Second answer should be 103, result was %d", second)
	}
	splittedStrings = Utilities.Read("example3.txt")
	second = Second(splittedStrings)
	if second != 3509 {
		t.Errorf("Second answer should be 3509, result was %d", second)
	}
}
