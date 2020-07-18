package algs

/*
FindSubstring solves the follwing problem:
	You are given a string, s, and a list of words, words, that are all of the same length.
	Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.


	Example 1:

		Input:
		s = "barfoothefoobarman",
		words = ["foo","bar"]
		Output: [0,9]
		Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
		The output order does not matter, returning [9,0] is fine too.

	Example 2:

		Input:
		s = "wordgoodgoodgoodbestword",
		words = ["word","good","best","word"]
		Output: []
*/
func FindSubstring(s string, words []string) []int {
	// 排除特殊情况
	if len(words) == 0 {
		return nil
	}
	wordLen := len(words[0])
	combLen := len(words) * wordLen
	if len(s) < combLen {
		return nil
	}

	var indices []int
	m := make(map[string]int)
	// 分wordLen批次比较子串
	for i := 0; i < wordLen; i++ {
		j := i
		for j+combLen <= len(s) {
			for _, word := range words {
				m[word]++
			}
			k := j
			step := wordLen
			for k+wordLen <= j+combLen {
				word := s[k : k+wordLen]
				count, ok := m[word]
				// 不相等则跳过
				if !ok {
					step = k - j + wordLen
					break
				}
				// 说明有重复等于words中某个元素的情况
				if count == 0 {
					break
				}
				m[word]--
				k += wordLen
			}
			if k == j+combLen {
				indices = append(indices, j)
			}
			j += step

			for _, word := range words {
				m[word] = 0
			}
		}
	}

	return indices
}
