package isPalindrome

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{"a", true},
	}
	for _, test := range tests {
		if result := isPalindrome(test.s); result != test.expected {
			t.Errorf(" %s got %t expected %t", test.s, result, test.expected)
		}
	}
}
