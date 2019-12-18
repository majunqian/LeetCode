package rotateArray

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedNums []int
	}{
		{"1", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"2", []int{-1}, 2, []int{-1}},
		{"3", []int{1, 2}, 5, []int{2, 1}},
		{"nil array", []int{}, 100, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rotate(test.nums, test.k)
			if !compareArray(test.nums, test.expectedNums) {
				t.Errorf("expect %v got %v", test.expectedNums, test.nums)
			}
		})
	}
}
func compareArray(a []int, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
