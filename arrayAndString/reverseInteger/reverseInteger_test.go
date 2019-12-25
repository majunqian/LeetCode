package reverseInteger

import (
	"math"
	"testing"
)

func TestReverseinteger(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{-100000, -1},
		{0, 0},
		{math.MaxInt32, 0},
		{math.MinInt32, 0},
	}
	for _, test := range tests {
		if re := reverse(test.x); re != test.expected {
			t.Errorf(" %d got %d expected %d", test.x, re, test.expected)
		}
	}
}
