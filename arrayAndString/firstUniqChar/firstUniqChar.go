package firstUniqChar

//Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
//
//Examples:
//
//s = "leetcode"
//return 0.
//
//s = "loveleetcode",
//return 2.
//Note: You may assume the string contain only lowercase letters.

func firstUniqChar(s string) int {
	dic := [26]int{}
	for i := range s {
		if dic[s[i]-'a'] < 2 {
			dic[s[i]-'a']++
		}
	}
	for i := range s {
		if dic[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
