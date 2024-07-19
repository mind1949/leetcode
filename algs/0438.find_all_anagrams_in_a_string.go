package algs

// FindAnagrams
//
// Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
// Example 1:
//
// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
//
// Example 2:
//
// Input: s = "abab", p = "ab"
// Output: [0,1,2]
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".
//
// Constraints:
//
//	1 <= s.length, p.length <= 3 * 104
//	s and p consist of lowercase English letters.
func FindAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	countP := countStr(p)
	countSub := countStr(s[:len(p)])
	var res []int
	i, j := 0, len(p)-1
	for i <= j {
		if countP == countSub {
			res = append(res, i)
		}
		i++
		j++
		if j >= len(s) {
			break
		}
		// 通过固定窗口大小驱动特点 减少不必要的数组初始化
		countSub[s[i-1]-'a'] -= 1
		countSub[s[j]-'a'] += 1
	}
	return res
}

func countStr(s string) [26]int {
	var count [26]int
	for _, char := range s {
		count[char-'a'] += 1
	}
	return count
}
