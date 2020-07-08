package algs

import "math"

/*
StrStr solves the following problem:
	Implement strStr().

	Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

	Example 1:

		Input: haystack = "hello", needle = "ll"
		Output: 2

	Example 2:

		Input: haystack = "aaaaa", needle = "bba"
		Output: -1

	Clarification:

	What should we return when needle is an empty string? This is a great question to ask during an interview.

	For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	head, tail := 0, len(needle)-1
	for tail < len(haystack) {
		if haystack[head:tail+1] == needle {
			return head
		}

		head++
		tail++
	}

	return -1
}

// StrStrRabinKarp 使用rabin-karp算法实现的StrStr
func StrStrRabinKarp(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	// 哈希算法
	h2Hash := func(i int) int {
		return int(haystack[i] - 'a')
	}
	n2Hash := func(i int) int {
		return int(needle[i] - 'a')
	}

	const (
		base   = 26
		module = 2 << 30
	)
	var hhash, nhash int
	for i := 0; i < len(needle); i++ {
		hhash = (hhash*base + h2Hash(i)) % module
		nhash = (nhash*base + n2Hash(i)) % module
	}
	if hhash == nhash {
		return 0
	}

	sub := int(math.Pow(float64(base), float64(len(needle)))) % module
	// 滚动哈希，判断字符串是否相等
	for i := 1; i < len(haystack)-len(needle)+1; i++ {
		hhash = (hhash*base - h2Hash(i-1)*sub + h2Hash(i+len(needle)-1)) % module
		if hhash == nhash {
			return i
		}
	}

	return -1
}
