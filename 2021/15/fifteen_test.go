package fifteen

import (
	"2021/Utilities"
	"testing"
)

func TestFirst(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	first := First(splittedStrings)
	if first != 40 {
		t.Errorf("First answer should be 40, result was %d", first)
	}
}

func TestSecond(t *testing.T) {
	splittedStrings := Utilities.Read("example.txt")
	second := Second(splittedStrings)
	if second != 315 {
		t.Errorf("Second answer should be 315, result was %d", second)
	}
}
