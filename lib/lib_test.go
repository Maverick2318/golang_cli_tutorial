package lib

import (
	"testing"
)

func TestSum(t *testing.T) {
	expect := 2
	got := Sum(1, 1)

	if got != expect {
		t.Fatalf("Expected %v got: %v", expect, got)
	}
}
