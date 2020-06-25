package algs

import "testing"

func TestRomanToInt(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	} {
		got := RomanToInt(c.input)
		if got != c.expect {
			t.Errorf("input: %q\t|\texpect: %d\t|\tgot: %d", c.input, c.expect, got)
		}
	}
}
