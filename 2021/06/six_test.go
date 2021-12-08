package six

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", ",")
	first := First(splittedStrings)
	if first != 5934 {
		t.Errorf("First answer should be 5934, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", ",")
	second := Second(splittedStrings)
	if second != 26984457539 {
		t.Errorf("Second answer should be 26984457539, result was %d", second)
	}
}
