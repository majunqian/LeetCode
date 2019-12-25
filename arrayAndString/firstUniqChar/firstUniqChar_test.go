package firstUniqChar

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"eeedfg", 3},
		{"b", 0},
		{"bbbbbb", -1},
		{"", -1},
		{"aabbc", 4},
	}
	for _, test := range tests {
		if result := firstUniqChar(test.s); result != test.expected {
			t.Errorf(" %s got %d expected %d", test.s, result, test.expected)
		}
	}
}
