package algs

import (
	"math"
	"sort"
)

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
	for _, word := range words {
		m[word]++
	}
	// 分wordLen批次比较子串
	for i := 0; i < wordLen; i++ {
		counter := make(map[string]int)
		l, r := i, i
		num := 0
		for r+wordLen <= len(s) {
			word := s[r : r+wordLen]
			r += wordLen
			// 若word在words中不存在
			if m[word] == 0 {
				l = r
				num = 0
				for word := range counter {
					counter[word] = 0
				}

				continue
			}

			counter[word]++
			num++
			// 若word重复了
			for counter[word] > m[word] {
				counter[s[l:l+wordLen]]--
				l += wordLen
				num--

				continue
			}
			// 若子串中的单词与words中的单词完全匹配
			if num == len(words) {
				indices = append(indices, l)
				counter[s[l:l+wordLen]]--
				num--
				l += wordLen
			}
		}
	}

	return indices
}

// FindSubstringWithHash 使用滚动哈希法实现FindSubString
func FindSubstringWithHash(s string, words []string) []int {
	// 排除特殊情况
	if len(words) == 0 {
		return nil
	}
	wordLen := len(words[0])
	combLen := wordLen * len(words)

	// 求出words组合的hash值wordsHash
	sort.Strings(words)
	m := make(map[string]int)
	for i := range words {
		m[words[i]] = i + 1
	}
	wordsHash := 0
	// FIXME:
	word2Hash := func(word string) int {
		return int(math.Pow(float64(len(words)), float64(m[word]))) % (1 << 61)
	}
	for i := 0; i < len(words); i++ {
		h := word2Hash(words[i])
		wordsHash += h
	}

	// 求出所有的合规的子串索引indcies
	var indices []int
	for i := 0; i < wordLen; i++ {
		l, r := i, i
		subStrHash := 0
		for l+combLen <= len(s) {
			for r-l+wordLen <= combLen {
				word := s[r : r+wordLen]
				h := word2Hash(word)
				subStrHash += h
				r += wordLen
			}
			if subStrHash == wordsHash {
				indices = append(indices, l)
			}

			h := word2Hash(s[l : l+wordLen])
			subStrHash -= h
			l += wordLen
		}
	}

	return indices
}
