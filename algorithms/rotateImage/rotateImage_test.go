package rotateImage

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3}},
		},
		{
			name: "",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16}},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rotate(test.matrix)
			if !comPareMatrix(test.matrix, test.expected) {
				t.Errorf("got %v expected %v", test.matrix, test.expected)
			}

		})
	}
}

func comPareMatrix(matrix [][]int, expected [][]int) bool {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != expected[i][j] {
				return false
			}
		}
	}
	return true
}
