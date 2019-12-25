// leetcode TwoSum
package twoSum

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
func TwoSum(nums []int, target int) []int {
	dic := make(map[int]int)
	for i, v := range nums {
		if value, ok := dic[target-v]; ok {
			return []int{value, i}
		}
		dic[v] = i
	}
	return []int{}
}
