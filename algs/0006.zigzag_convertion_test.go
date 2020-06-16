package algs

import "testing"

func TestConvert(t *testing.T) {
	cs := []struct {
		s       string
		numRows int
		expect  string
	}{
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			expect:  "PAHNAPLSIIGYIR",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
			expect:  "PINALSIGYAHRPI",
		},
		{
			s:       "",
			numRows: 4,
			expect:  "",
		},
		{
			s:       "AB",
			numRows: 1,
			expect:  "AB",
		},
	}

	for _, c := range cs {
		got := Convert(c.s, c.numRows)
		if got != c.expect {
			t.Errorf("s: %q\t|\tnumRows: '%d'\t|\texpect: %q\t|\tgot: %q", c.s, c.numRows, c.expect, got)
		}
	}
}
