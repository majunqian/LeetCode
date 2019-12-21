package singleNumber

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		expectedResult int
	}{
		{"",[]int{2,2,1},1},
		{"",[]int{4,1,2,1,2},4},
		{"",[]int{4},4},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			number := singleNumber(test.nums)
			if number != test.expectedResult{
				t.Errorf("%v expect %d got %d",
					test.nums, test.expectedResult, number)
			}
		})
	}
}
