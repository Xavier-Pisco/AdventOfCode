package fourteen

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", "\n\n")
	first := First(splittedStrings)
	if first != 1588 {
		t.Errorf("First answer should be 1588, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", "\n\n")
	second := Second(splittedStrings)
	if second != 2188189693529 {
		t.Errorf("Second answer should be 2188189693529, result was %d", second)
	}
}
