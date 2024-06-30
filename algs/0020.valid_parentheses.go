package algs

/*
IsValid solves the following problem:

	Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

	An input string is valid if:

		Open brackets must be closed by the same type of brackets.
		Open brackets must be closed in the correct order.

	Note that an empty string is also considered valid.

	Example 1:

		Input: "()"
		Output: true

	Example 2:

		Input: "()[]{}"
		Output: true

	Example 3:

		Input: "(]"
		Output: false

	Example 4:

		Input: "([)]"
		Output: false

	Example 5:

		Input: "{[]}"
		Output: true
*/
func IsValid(s string) bool {
	// 括号对
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	lps := make([]byte, 0, len(s)/2) // 左括号
	for i := 0; i < len(s); i++ {
		p := s[i]
		if _, ok := m[p]; ok {
			lps = append(lps, p)
		} else {
			if len(lps) > 0 && m[lps[len(lps)-1]] == p {
				lps = lps[:len(lps)-1]
			} else {
				return false
			}
		}
	}

	if len(lps) == 0 {
		return true
	}
	return false
}
