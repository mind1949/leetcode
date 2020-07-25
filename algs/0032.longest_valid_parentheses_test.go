package algs

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	for _, c := range []struct {
		s      string
		expect int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"((((", 0},
		{"))))", 0},
		{")))(((", 0},
	} {
		got := LongestValidParentheses(c.s)
		if got != c.expect {
			t.Errorf("s: %q\t|\texpect: %d\t|\tgot: %d", c.s, c.expect, got)
		}
	}
}
