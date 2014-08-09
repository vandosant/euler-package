package euler

import (
	"testing"
)

func TestSmallestMultiple(t *testing.T) {
	r := SmallestMultiple(10)
	if r != 2520 {
		t.Error("Test failed.")
	}
}
