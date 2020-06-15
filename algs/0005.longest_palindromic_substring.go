package algs

/*
LongestPalindrome soloves following problem:
	Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

	Example 1:

		Input: "babad"
		Output: "bab"
		Note: "aba" is also a valid answer.

	Example 2:

		Input: "cbbd"
		Output: "bb"
*/
func LongestPalindrome(s string) string {
	list := []rune(s)
	if len(list) <= 1 {
		return s
	}

	var max = []rune{list[0]}
	for i := 0; i < len(list); i++ {
		// case one: len(subList)%2 == 1
		left, right := i-1, i+1
		for left >= 0 && right < len(list) {
			if list[right] != list[left] {
				break
			}
			if right-left+1 > len(max) {
				max = list[left : right+1]
			}
			left, right = left-1, right+1
		}
		// case two: len(subList)%2 === 0
		left, right = i-1, i
		for left >= 0 && right < len(list) {
			if list[right] != list[left] {
				break
			}
			if right-left+1 > len(max) {
				max = list[left : right+1]
			}
			left, right = left-1, right+1
		}
	}

	return string(max)
}
