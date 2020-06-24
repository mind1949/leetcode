package algs

import "testing"

func TestIntToRoman(t *testing.T) {
	for _, c := range []struct {
		input  int
		expect string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	} {
		got := IntToRoman(c.input)
		if got != c.expect {
			t.Errorf("input: %d \t|\texpect: %q\t|\tgot: %q", c.input, c.expect, got)
		}
	}
}
