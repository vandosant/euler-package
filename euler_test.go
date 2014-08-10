package euler

import (
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	a := SmallestMultiple(10)
	e := 2520
	if a != 2520 {
		t.Error("Test failed. Expected:", e, "Actual:", a)
	}
}

func TestSumSquareDifference(t *testing.T) {
	a := SumSquareDifference(10)
	e := 2640
	if a != e {
		t.Error("Test failed. Expected:", e, "Actual:", a)
	}
}

func TestNthPrimeNumber(t *testing.T) {
	a := NthPrimeNumber(6)
	e := 13
	if a != e {
		t.Error("Test failed. Expected:", e, "Actual:", a)
	}
}
