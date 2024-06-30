package algs

/*
LongestValidParentheses solves the following problem:

	Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

	Example 1:

	Input: "(()"
	Output: 2
	Explanation: The longest valid parentheses substring is "()"

	Example 2:

	Input: ")()())"
	Output: 4
	Explanation: The longest valid parentheses substring is "()()"
*/
func LongestValidParentheses(s string) int {
	// 排除不可能有合法括号的字符串
	bs := []byte(s)
	i := 0
	for i < len(bs) && bs[i] == ')' { // 去除左边不可能配对成功的
		i++
	}
	bs = bs[i:]

	i = len(bs) - 1
	for i >= 0 && bs[i] == '(' { // 去除右边不可能配对成功的
		i--
	}
	bs = bs[0 : i+1]

	if len(bs) < 2 {
		return 0
	}

	// 第一步遍历标记所有配对的位置
	i = 0
	for i < len(bs) {
		if bs[i] == ')' {
			j := i - 1
			for j >= 0 && bs[j] == '|' {
				j--
			}
			if j >= 0 && bs[j] == '(' {
				bs[j] = '|'
				bs[i] = '|'
			}
		}
		i++
	}

	// 第二步计算最长的括号长度
	max := 0
	cur := 0
	for i := 0; i < len(bs); i++ {
		if bs[i] == '|' {
			cur++
		} else {
			if cur > max {
				max = cur
			}
			cur = 0
		}
	}
	if cur > max {
		max = cur
	}

	return max
}

// LongestValidParentheses1 左右迭代对称求值
func LongestValidParentheses1(s string) int {
	var (
		l, r, maxLen int
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			maxLen = max(maxLen, 2*r)
		} else if r > l {
			l, r = 0, 0
		}
	}

	r, l = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			r++
		} else {
			l++
		}
		if l == r {
			maxLen = max(maxLen, 2*l)
		} else if l > r {
			l, r = 0, 0
		}
	}
	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
