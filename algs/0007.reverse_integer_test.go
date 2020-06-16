package algs

import "testing"

func TestReverse(t *testing.T) {
	cs := []struct {
		input  int
		expect int
	}{
		{
			input:  321,
			expect: 123,
		},
		{
			input:  120,
			expect: 21,
		},
		{
			input:  102,
			expect: 201,
		},
		{
			input:  -123,
			expect: -321,
		},
		{
			input:  1534236469,
			expect: 0,
		},
		{
			input:  32768,
			expect: 86723,
		},
	}

	for _, c := range cs {
		got := Reverse(c.input)
		if got != c.expect {
			t.Errorf("input: '%d'\t|\texpect: '%d'\t|\tgot: '%d'", c.input, c.expect, got)
		}
	}
}
