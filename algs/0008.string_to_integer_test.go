package algs

import "testing"

func TestMyAtoi(t *testing.T) {
	cs := []struct {
		input  string
		expect int
	}{
		{
			input:  "42",
			expect: 42,
		},
		{
			input:  "   -42",
			expect: -42,
		},
		{
			input:  "4193 with words",
			expect: 4193,
		},
		{
			input:  "words and 987",
			expect: 0,
		},
		{
			input:  "-91283472332",
			expect: -2147483648,
		},
		{
			input:  "-01",
			expect: -1,
		},
		{
			input:  "-0",
			expect: 0,
		},
		{
			input:  "-",
			expect: 0,
		},
		{
			input:  "+",
			expect: 0,
		},
		{
			input:  "9223372036854775808",
			expect: 2147483647,
		},
	}

	for _, c := range cs {
		got := MyAtoi(c.input)
		if got != c.expect {
			t.Errorf("input: %s\t|\texpect: %d\t|\tgot: %d", c.input, c.expect, got)
		}
	}
}
