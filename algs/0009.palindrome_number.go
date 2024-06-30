package algs

import (
	"math"
)

/*
IsPalindrome solves the following problem:

	Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

	Example 1:

		Input: 121
		Output: true

	Example 2:

		Input: -121
		Output: false
		Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

	Example 3:

		Input: 10
		Output: false
		Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

	Follow up:

		Coud you solve it without converting the integer to a string?
*/
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	// 先求出总共有多少位
	n := 2
	for {
		if x/int(math.Pow10(n)) < 1 {
			break
		}
		n++
	}
	// 比较对称位上的值是否相等
	for i, j := 1, n; i <= n/2; i, j = i+1, n+1-i {
		r := (x % int(math.Pow10(i))) / int(math.Pow10(i-1))
		l := (x % int(math.Pow10(j))) / int(math.Pow10(j-1))

		if r != l {
			return false
		}
	}

	return true
}
