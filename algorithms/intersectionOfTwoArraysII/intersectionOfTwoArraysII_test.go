package intersectionOfTwoArraysII

import (
	"sort"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name           string
		nums1          []int
		nums2          []int
		expectedResult []int
	}{
		{"", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{"", []int{1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ints := intersect(test.nums1, test.nums2)
			sort.Ints(ints)
			if len(ints) != len(test.expectedResult) &&
				compareArray(test.nums1, test.nums2) {
				t.Errorf("%v %v expected %v got %v ",
					test.nums1, test.nums2, test.expectedResult, ints)
			}
		})
	}
}
func compareArray(nums1 []int, nums2 []int) bool {
	for i := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
