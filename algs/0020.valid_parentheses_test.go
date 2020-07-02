package algs

import "testing"

func TestIsValid(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect bool
	}{
		{"", true},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"{[]}", true},
	} {
		got := IsValid(c.input)
		if got != c.expect {
			t.Errorf("input: %q\t|\texpect: %t\t|\tgot: %t", c.input, c.expect, got)
		}
	}
}
