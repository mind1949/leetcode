package algs

/*
Divide soloves the following problem:

	Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

	Return the quotient after dividing dividend by divisor.

	The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.

	Example 1:

		Input: dividend = 10, divisor = 3
		Output: 3
		Explanation: 10/3 = truncate(3.33333..) = 3.

	Example 2:

		Input: dividend = 7, divisor = -3
		Output: -2
		Explanation: 7/-3 = truncate(-2.33333..) = -2.

	Note:

		Both dividend and divisor will be 32-bit signed integers.
		The divisor will never be 0.
		Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 2^31 − 1 when the division result overflows.
*/
func Divide(dividend int, divisor int) int {
	// 判断被除数与除数的正负性
	positive := true
	if dividend < 0 {
		positive = !positive
		dividend = -dividend
	}
	if divisor < 0 {
		positive = !positive
		divisor = -divisor
	}

	// 计算商
	quotient := 0
	if divisor == 1 {
		quotient = dividend
		dividend = 0
	}
	guess, compare := 31, 0
	for compare+divisor <= dividend {
		if compare+divisor<<guess <= dividend {
			compare += divisor << guess
			quotient += 1 << guess
		}
		guess--
	}

	if !positive {
		quotient = -quotient
	}
	if quotient > 1<<31-1 {
		quotient = 1<<31 - 1
	}
	return quotient
}
