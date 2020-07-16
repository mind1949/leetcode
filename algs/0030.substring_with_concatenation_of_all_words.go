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

	// 求words的所有组合
	var (
		combs   = []string{""}                     // words所有的组合
		counter = make(map[string]int, len(words)) // 每一次组合中可用的word数量
	)
	for i := 0; i < len(words); i++ {
		combsLen := len(combs)
		for j := 0; j < combsLen; j++ {
			comb := combs[j]
			for _, word := range words {
				counter[word]++
			}
			for head := 0; head+wordLen <= len(comb); head += wordLen {
				word := comb[head : head+wordLen]
				counter[word]--
			}

			k := 0
			for _, word := range words {
				if counter[word] > 0 {
					if k == 0 {
						combs[j] = comb + word
					} else {
						combs = append(combs, comb+word)
					}
					k++
					counter[word]--
				}
			}
		}

	}

	// 遍历比较所有的子字符串
	var indices []int
	for head := 0; head+combLen <= len(s); head++ {
		subStr := s[head : head+combLen]
		for _, comb := range combs {
			if comb == subStr {
				indices = append(indices, head)
				break
			}
		}
	}

	return indices
}
