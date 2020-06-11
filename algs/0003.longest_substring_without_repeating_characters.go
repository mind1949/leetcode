package algs

/*
LengthOfLongestSubstring solves following problem:
	Given a string, find the length of the longest substring without repeating characters.

	Example 1:

		Input: "abcabcbb"
		Output: 3
		Explanation: The answer is "abc", with the length of 3.

	Example 2:

		Input: "bbbbb"
		Output: 1
		Explanation: The answer is "b", with the length of 1.

	Example 3:

		Input: "pwwkew"
		Output: 3
		Explanation: The answer is "wke", with the length of 3.
					Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func LengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	maxLen := 0
	start := 0
	m := make(map[rune]int, len(runes))
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if k, ok := m[r]; ok {
			if k >= start {
				l := i - start
				if maxLen < l {
					maxLen = l
				}
				start = k + 1
			}
		}
		m[r] = i
	}
	l := len(runes) - start
	if maxLen < l {
		maxLen = l
	}

	return maxLen
}
