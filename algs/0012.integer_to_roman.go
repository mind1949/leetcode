package algs

import (
	"fmt"
	"math"
)

/*
IntToRoman solves the following problem:

	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

		Symbol       Value
		I             1
		V             5
		X             10
		L             50
		C             100
		D             500
		M             1000

	For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

	Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

		I can be placed before V (5) and X (10) to make 4 and 9.
		X can be placed before L (50) and C (100) to make 40 and 90.
		C can be placed before D (500) and M (1000) to make 400 and 900.

	Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

	Example 1:

		Input: 3
		Output: "III"

	Example 2:

		Input: 4
		Output: "IV"

	Example 3:

		Input: 9
		Output: "IX"

	Example 4:

		Input: 58
		Output: "LVIII"
		Explanation: L = 50, V = 5, III = 3.

	Example 5:

		Input: 1994
		Output: "MCMXCIV"
		Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/
func IntToRoman(num int) (r string) {
	if num > 3999 || num < 1 {
		return ""
	}

	// 基本元素
	e := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	// 从个位往上开始逐个拼凑数值
	for i := 0; num >= 1; i++ {
		v := num % 10           // 看它是几
		g := int(math.Pow10(i)) // 它属于第几位
		switch v {
		case 0:
			// no-op
		case 1, 2, 3:
			for i := 0; i < v; i++ {
				r = e[g] + r
			}
		case 4:
			r = e[g] + e[5*g] + r
		case 5:
			r = e[5*g] + r
		case 6, 7, 8:
			for i := 0; i < v-5; i++ {
				r = e[g] + r
			}
			r = e[5*g] + r
		case 9:
			r = e[g] + e[g*10] + r
		default:
			panic(fmt.Sprintf("impossible: %d", v))
		}
		num /= 10
	}

	return r
}
