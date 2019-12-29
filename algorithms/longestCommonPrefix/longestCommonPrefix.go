package longestCommonPrefix

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//Example 1:
//
//Input: ["flower","flow","flight"]
//Output: "fl"
//Example 2:
//
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
//Note:
//
//All given inputs are in lowercase letters a-z.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i <= len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if !hasPrefix(strs[j], strs[0][0:i]) {
				return strs[0][0 : i-1]
			}
		}
	}
	return strs[0]
}

// strings.HasPrefix
func hasPrefix(str string, prefix string) bool {
	return len(str) >= len(prefix) && str[0:len(prefix)] == prefix
}
