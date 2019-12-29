package longestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{"", []string{"flower", "flow", "flight"}, "fl"},
		{"", []string{"dog", "racecar", "car"}, ""},
		{"", []string{"dog", "", "car"}, ""},
		{"", []string{"flower", "flow", "fhight"}, "f"},
		{"", []string{"", "flow", "fhight"}, ""},
		{"", []string{"a"}, "a"},
		{"", []string{}, ""},
		{"", []string{"a", "b"}, ""},
		{"", []string{"a", "a", "b"}, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if re := longestCommonPrefix(test.strs); re != test.expected {
				t.Errorf("%v got %s expected %s",
					test.strs, re, test.expected)
			}
		})
	}
}
