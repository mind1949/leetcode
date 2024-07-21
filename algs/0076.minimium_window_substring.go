package algs

// MinWindow
//
// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring
// of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
//
// The testcases will be generated such that the answer is unique.
//
// Example 1:
//
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
//
// Example 2:
//
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.
//
// Example 3:
//
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.
//
// Constraints:
//
//	m == s.length
//	n == t.length
//	1 <= m, n <= 105
//	s and t consist of uppercase and lowercase English letters.
//
// Follow up: Could you find an algorithm that runs in O(m + n) time?
func MinWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	var counter [58]int
	get := func(s string, i int) byte {
		return s[i] - byte('A')
	}
	for i := 0; i < len(t); i++ {
		counter[get(t, i)] += 1
	}

	var counterSub [58]int
	counterSub[get(s, 0)] += 1
	contain := func(src, target [58]int) bool {
		for i := 0; i < len(target); i++ {
			if src[i] >= target[i] {
				continue
			}
			return false
		}
		return true
	}
	left, right := 0, 0
	i, j := 0, 0
	for i <= j && j < len(s) {
		if contain(counterSub, counter) {
			if right-left == 0 || j+1-i < right-left {
				left, right = i, j+1
			}
			counterSub[get(s, i)] -= 1
			i++
		} else {
			j++
			if j < len(s) {
				counterSub[get(s, j)] += 1
			}
		}
	}
	return s[left:right]
}
