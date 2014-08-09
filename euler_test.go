package euler

import (
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	a := SmallestMultiple(10)
	if a != 2520 {
		t.Error("Test failed.")
	}
}

func TestSumSquareDifference(t *testing.T) {
	a := SumSquareDifference(10)
	e := 2640
	if a != e {
		t.Error("Test failed. Expected:", e, "Actual:", a)
	}
}
