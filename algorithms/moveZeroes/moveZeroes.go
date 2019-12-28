package moveZeroes

//Given an arrayAndString nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//Example:
//
//Input: [0,1,0,3,12]
//Output: [1,3,12,0,0]
//Note:
//
//You must do this in-place without making a copy of the arrayAndString.
//Minimize the total number of operations

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	var j int
	for i := range nums {
		if nums[i] != 0 {
			if i != j {
				nums[j] = nums[i]
			}
			j++
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
