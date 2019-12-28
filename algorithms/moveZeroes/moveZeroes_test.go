package moveZeroes

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"", []int{0}, []int{0}},
		{"", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"", []int{1, 2, 3, 4, 0, 0}, []int{1, 2, 3, 4, 0, 0}},
		{"", []int{}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			moveZeroes(test.nums)
			for i := range test.nums {
				if test.nums[i] != test.expected[i] {
					t.Errorf("got %v expected %v",
						test.nums, test.expected)
					break
				}
			}
		})
	}
}
