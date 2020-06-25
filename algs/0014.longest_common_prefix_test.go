package algs

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	for _, c := range []struct {
		input  []string
		expect string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{nil, ""},
		{[]string{"flower", "flow", "flowers"}, "flow"},
	} {
		got := LongestCommonPrefix(c.input)
		if got != c.expect {
			t.Errorf("input: %v\t|\texpect: %q\t|\tgot: %q", c.input, c.expect, got)
		}
	}
}
