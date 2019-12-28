package containsDuplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		expectedResult bool
	}{
		{"", []int{1, 2, 3, 1}, true},
		{"", []int{1, 2, 3, 4}, false},
		{"", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			duplicate := containsDuplicate(test.nums)
			if duplicate != test.expectedResult {
				t.Errorf("%v expected %v got %v",
					test.nums, test.expectedResult, duplicate)
			}
		})
	}
}
