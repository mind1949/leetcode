package algs

// GroupAnagrams solves the follow problem:
//
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
// Example 1:
//
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Example 2:
//
// Input: strs = [""]
// Output: [[""]]
//
// Example 3:
//
// Input: strs = ["a"]
// Output: [["a"]]
//
// Constraints:
//
//	1 <= strs.length <= 104
//	0 <= strs[i].length <= 100
//	strs[i] consists of lowercase English letters.
func GroupAnagrams(strs []string) [][]string {
	matchers := make(map[[26]int]int)
	result := [][]string{}
	for _, str := range strs {
		matcher := [26]int{}
		for _, char := range str {
			matcher[char-'a']++
		}
		idx, ok := matchers[matcher]
		if ok {
			result[idx] = append(result[idx], str)
		} else {
			matchers[matcher] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}
