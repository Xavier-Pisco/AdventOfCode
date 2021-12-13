package thirteen

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", "\n\n")
	first := First(splittedStrings)
	if first != 17 {
		t.Errorf("First answer should be 17, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", "\n\n")
	second := Second(splittedStrings)
	if second != 16 {
		t.Errorf("Second answer should be 16, result was %d", second)
	}
}
