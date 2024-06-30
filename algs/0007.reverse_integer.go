package algs

import (
	"strconv"
)

/*
Reverse solves the following problem:

	Given a 32-bit signed integer, reverse digits of an integer.

	Example 1:

		Input: 123
		Output: 321

	Example 2:

		Input: -123
		Output: -321

	Example 3:

		Input: 120
		Output: 21

	Note:
		Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
func Reverse(x int) int {
	positive := 1
	if x < 0 {
		positive = -1
		x = -x
	}

	var s string
	for _, v := range strconv.Itoa(x) {
		s = string(v) + s
	}
	v, _ := strconv.Atoi(s)
	v = v * positive

	if int(int32(v)) != v {
		v = 0
	}
	return v
}
