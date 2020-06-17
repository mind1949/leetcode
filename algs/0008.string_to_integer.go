package algs

import (
	"math"
	"unicode"
)

/*
MyAtoi solves the following problem:
	Implement atoi which converts a string to an integer.

	The function first discards as many whitespace characters as necessary
	until the first non-whitespace character is found.
	Then, starting from this character, takes an optional initial plus or
	minus sign followed by as many numerical digits as possible,
	and interprets them as a numerical value.

	The string can contain additional characters after those
	that form the integral number, which are ignored and have no effect on the behavior of this function.

	If the first sequence of non-whitespace characters in str is not a valid integral number,
	or if no such sequence exists because either str is empty or
	it contains only whitespace characters, no conversion is performed.

	If no valid conversion could be performed, a zero value is returned.

	Note:

		Only the space character ' ' is considered as whitespace character.
		Assume we are dealing with an environment which could only store integers
		within the 32-bit signed integer range: [−2^31,  2^31 − 1].
		If the numerical value is out of the range of representable values, INT_MAX (2^31 − 1) or INT_MIN (−2^31) is returned.

	Example 1:

		Input: "42"
		Output: 42

	Example 2:

		Input: "   -42"
		Output: -42
		Explanation: The first non-whitespace character is '-', which is the minus sign.
					Then take as many numerical digits as possible, which gets 42.

	Example 3:

		Input: "4193 with words"
		Output: 4193
		Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.

	Example 4:

		Input: "words and 987"
		Output: 0
		Explanation: The first non-whitespace character is 'w', which is not a numerical
					digit or a +/- sign. Therefore no valid conversion could be performed.

	Example 5:

		Input: "-91283472332"
		Output: -2147483648
		Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
					Thefore INT_MIN (−2^31) is returned.
*/
func MyAtoi(str string) int {
	var nums []rune
	i := 0
	for _, v := range str {
		// 过滤掉第一个数字或者正负号前的空格
		if v == rune(' ') && i == 0 {
			continue
		}
		// 判断合法性
		if i == 0 && !unicode.IsNumber(v) && v != rune('-') && v != rune('+') {
			return 0
		}
		// 过滤掉数字后面的非数字
		if i > 0 && !unicode.IsNumber(v) {
			break
		}
		nums = append(nums, v)
		i++
	}
	if len(nums) == 0 {
		return 0
	}

	// 计算
	op := 1
	if nums[0] == rune('-') {
		op = -1
		nums = nums[1:]
	} else if nums[0] == rune('+') {
		nums = nums[1:]
	}

	num := 0
	for _, v := range nums {
		num = num*10 + int(v-rune('0'))
		// 提前结束，避免结果大于int64时溢出，导致结果不正确
		if num > -int(math.MinInt32) {
			break
		}
	}
	num *= op

	// 判断值是否超过了int32的范围
	if num > math.MaxInt32 {
		num = math.MaxInt32
	}
	if num < math.MinInt32 {
		num = math.MinInt32
	}

	return num
}
