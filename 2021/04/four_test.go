package four

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", "\n\n")
	first := First(splittedStrings)
	if first != 4512 {
		t.Errorf("First answer should be 4512, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", "\n\n")
	second := Second(splittedStrings)
	if second != 1924 {
		t.Errorf("Second answer should be 1924, result was %d", second)
	}
}
