package algs

/*
LongestCommonPrefix solves the following problem:
	Write a function to find the longest common prefix string amongst an array of strings.

	If there is no common prefix, return an empty string "".

	Example 1:

		Input: ["flower","flow","flight"]
		Output: "fl"

	Example 2:

	Input: ["dog","racecar","car"]
		Output: ""
		Explanation: There is no common prefix among the input strings.

	Note:

	All given inputs are in lowercase letters a-z.
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		const any byte = '*'
		var pre = any // 哨兵
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) == i {
				return strs[j]
			}
			cur := strs[j][i]
			if cur != pre && pre != any {
				return strs[j][:i]
			}
			pre = cur
		}
	}
}
