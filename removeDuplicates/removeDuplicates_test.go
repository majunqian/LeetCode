package removeDuplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name        string
		array       []int
		expectedLen int
	}{
		{"normal", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{"normal", []int{0, 1, 2, 3, 4, 5, 6}, 7},
		{"one element", []int{0, 0, 0}, 1},
		{"nil array", []int{}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			length := removeDuplicates(test.array)
			if test.expectedLen != length {
				t.Errorf("Array %v expected length %d got %d",
					test.array, test.expectedLen, length)
			}

		})
	}
}
