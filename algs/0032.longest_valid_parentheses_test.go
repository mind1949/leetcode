package algs

import "testing"

var cases0032 = []struct {
	s      string
	expect int
}{
	{"(()", 2},
	{")()())", 4},
	{"", 0},
	{"((((", 0},
	{"))))", 0},
	{")))(((", 0},
}

func TestLongestValidParentheses(t *testing.T) {
	for _, c := range cases0032 {
		got := LongestValidParentheses(c.s)
		if got != c.expect {
			t.Errorf("s: %q\t|\texpect: %d\t|\tgot: %d", c.s, c.expect, got)
		}
	}
}

func TestLongestValidParenthesesSymmetry(t *testing.T) {
	for _, c := range cases0032 {
		got := LongestValidParentheses1(c.s)
		if got != c.expect {
			t.Errorf("s: %q\t|\texpect: %d\t|\tgot: %d", c.s, c.expect, got)
		}
	}
}
