package algs

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, c := range []struct {
		input  int
		expect bool
	}{
		{
			input:  121,
			expect: true,
		},
		{
			input:  -121,
			expect: false,
		},
		{
			input:  231,
			expect: false,
		},
		{
			input:  3,
			expect: true,
		},
	} {
		got := IsPalindrome(c.input)
		if got != c.expect {
			t.Errorf("input: %d\t|\texpect: %t\t|\tgot: %t", c.input, c.expect, got)
		}
	}
}
