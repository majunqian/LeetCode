package strStr

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{"", "hello", "ll", 2},
		{"", "aaaaa", "bba", -1},
		{"", "aaaaa", "", 0},
		{"", "aaa", "a", 0},
		{"", "abcdeef", "ef", 5},
		{"", "abcdeef", "fe", -1},
		{"", "abcdeef", "abcdeef", 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := strStr(test.haystack, test.needle); result != test.expected {
				t.Errorf("%s %s got %d expected %d",
					test.haystack, test.needle, result, test.expected)
			}
		})
	}
}
