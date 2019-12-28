package myAtoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected int
	}{
		{"0", "42", 42},
		{"1", "-42", -42},
		{"2", "4193 with words", 4193},
		{"3", "words and 987", 0},
		{"4", "-91283472332", -2147483648},
		{"5", "9223372036854775808", 2147483647},
		{"6", "3.1415926", 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := myAtoi(test.str); result != test.expected {
				t.Errorf("%s got %d expected %d",
					test.str, result, test.expected)
			}
		})
	}
}
