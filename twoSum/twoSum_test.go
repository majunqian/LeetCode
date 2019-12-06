package twoSum

import (
	"fmt"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	test := []struct {
		numbers []int
		target  int
		result  []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{5, 5}, 10, []int{0, 1}},
	}
	for _, tt := range test {
		actual := TwoSum(tt.numbers, tt.target)
		sort.Ints(actual)
		for i, get := range actual {
			if get != tt.result[i] {
				t.Errorf("numbers %v target %d \nexpect %v get %v\n",
					tt.numbers, tt.target, tt.result, actual)
			}
		}
	}
}

func ExampleTwoSum() {
	nums := []int{1, 2, 3, 4, 5}
	result := TwoSum(nums, 9)
	fmt.Println(result)

	//OUTPUT:
	//[3 4]
}
