package reverseString

//Write a function that reverses a string. The input string is given as an arrayAndString of characters char[].
//
//Do not allocate extra space for another arrayAndString, you must do this by modifying the input arrayAndString in-place with O(1) extra memory.
//
//You may assume all the characters consist of printable ascii characters.
//
//
//
//Example 1:
//
//Input: ["h","e","l","l","o"]
//Output: ["o","l","l","e","h"]
//Example 2:
//
//Input: ["H","a","n","n","a","h"]
//Output: ["h","a","n","n","a","H"]

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
