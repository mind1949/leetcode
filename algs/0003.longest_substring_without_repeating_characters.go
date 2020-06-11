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
	lgs := 0
	for i := 0; i < len(runes); i++ {
		m := make(map[rune]struct{}, len(runes))
		for j := i; j < len(runes); j++ {
			if _, ok := m[runes[j]]; ok {
				break
			}
			m[runes[j]] = struct{}{}
		}
		c := len(m)
		if lgs < c {
			lgs = c
		}
	}

	return lgs
}
