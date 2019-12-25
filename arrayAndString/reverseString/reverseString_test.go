package reverseString

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		s        []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte(""), []byte("")},
		{[]byte("1"), []byte("1")},
		{[]byte("12345678900---"), []byte("---00987654321")},
	}
	for _, test := range tests {
		reverseString(test.s)
		if string(test.s) != string(test.expected) {
			t.Errorf("Got %s expected %s", test.s, test.expected)
		}

	}
}
