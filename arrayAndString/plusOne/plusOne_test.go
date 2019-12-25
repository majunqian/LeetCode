package plusOne

import "testing"

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"", []int{1, 2, 3}, []int{1, 2, 4}},
		{"", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"", []int{0}, []int{1}},
		{"", []int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := plusOne(test.nums)
			if len(result) != len(test.expected) {
				t.Errorf("%v  expect %v got %v",
					test.nums, test.expected, result)
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("%v  expect %v got %v",
						test.nums, test.expected, result)
					break
				}
			}
		})
	}
}
