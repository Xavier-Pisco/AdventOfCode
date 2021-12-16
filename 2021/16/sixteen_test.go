package sixteen

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	first := First(splittedStrings)
	if first != 16 {
		t.Errorf("First 1 answer should be 16, result was %d", first)
	}
	splittedStrings = Utilities.Read("example2.txt")
	first = First(splittedStrings)
	if first != 12 {
		t.Errorf("First 2 answer should be 12, result was %d", first)
	}
	splittedStrings = Utilities.Read("example3.txt")
	first = First(splittedStrings)
	if first != 23 {
		t.Errorf("First 3 answer should be 23, result was %d", first)
	}
	splittedStrings = Utilities.Read("example4.txt")
	first = First(splittedStrings)
	if first != 31 {
		t.Errorf("First 4 answer should be 31, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.Read("example5.txt")
	second := Second(splittedStrings)
	if second != 3 {
		t.Errorf("Second 5 answer should be 3, result was %d", second)
	}
	splittedStrings = Utilities.Read("example6.txt")
	second = Second(splittedStrings)
	if second != 54 {
		t.Errorf("Second 6 answer should be 54, result was %d", second)
	}
	splittedStrings = Utilities.Read("example7.txt")
	second = Second(splittedStrings)
	if second != 7 {
		t.Errorf("Second 7 answer should be 7, result was %d", second)
	}
	splittedStrings = Utilities.Read("example8.txt")
	second = Second(splittedStrings)
	if second != 9 {
		t.Errorf("Second 8 answer should be 9, result was %d", second)
	}
	splittedStrings = Utilities.Read("example9.txt")
	second = Second(splittedStrings)
	if second != 1 {
		t.Errorf("Second 9 answer should be 1, result was %d", second)
	}
	splittedStrings = Utilities.Read("example10.txt")
	second = Second(splittedStrings)
	if second != 0 {
		t.Errorf("Second 10 answer should be 0, result was %d", second)
	}
	splittedStrings = Utilities.Read("example11.txt")
	second = Second(splittedStrings)
	if second != 0 {
		t.Errorf("Second 11 answer should be 0, result was %d", second)
	}
	splittedStrings = Utilities.Read("example12.txt")
	second = Second(splittedStrings)
	if second != 1 {
		t.Errorf("Second 12 answer should be 1, result was %d", second)
	}
	// real = 1495959086337
}
