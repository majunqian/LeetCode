package removeDuplicates

//Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
//
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
//
//Example 1:
//
//Given nums = [1,1,2],
//
//Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
//
//It doesn't matter what you leave beyond the returned length

func removeDuplicates(nums []int) int {
	arrayLen := len(nums)
	if arrayLen < 2 {
		return arrayLen
	}
	pos := 0
	for i := 1; i < arrayLen; i++ {
		if nums[i] > nums[pos] {
			nums[pos+1] = nums[i]
			pos++
		}
	}
	return pos + 1
}