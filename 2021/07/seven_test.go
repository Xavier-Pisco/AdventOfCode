package seven

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", ",")
	first := First(splittedStrings)
	if first != 37 {
		t.Errorf("First answer should be 37, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.ReadWithDelimiter("example.txt", ",")
	second := Second(splittedStrings)
	if second != 168 {
		t.Errorf("Second answer should be 168, result was %d", second)
	}
}
