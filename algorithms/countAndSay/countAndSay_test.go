package countAndSay

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected string
	}{
		{"", 1, "1"},
		{"", 2, "11"},
		{"", 3, "21"},
		{"", 4, "1211"},
		{"", 5, "111221"},
		{"", 6, "312211"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := countAndSay(test.num); result != test.expected {
				t.Errorf(" %d got %s expected %s", test.num, result, test.expected)
			}
		})
	}
}
