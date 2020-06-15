package algs

import (
	"testing"
)

/*
Example 1:

		Input: "babad"
		Output: "bab"
		Note: "aba" is also a valid answer.

	Example 2:

		Input: "cbbd"
		Output: "bb"
*/
func TestLongestPalindrome(t *testing.T) {
	cs := []struct {
		input  string
		expect string
	}{
		{
			input:  "babad",
			expect: "bab",
		},
		{
			input:  "cbbd",
			expect: "bb",
		},
		{
			input:  "a",
			expect: "a",
		},
		{
			input:  "",
			expect: "",
		},
	}

	for _, c := range cs {
		got := LongestPalindrome(c.input)
		if got != c.expect {
			t.Errorf("input: %q \t|\texpect: %q\t|\tgot: %q", c.input, c.expect, got)
		}
	}
}
