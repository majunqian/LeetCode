package isAnagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"123中国我爱你", "我爱你中国123", true},
	}
	for _, test := range tests {
		if result := isAnagram(test.s, test.t); result != test.expected {
			t.Errorf(" %s %s got %t expected %t",
				test.s, test.t, result, test.expected)
		}
	}
}
