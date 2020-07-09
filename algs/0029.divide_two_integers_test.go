package algs

import (
	"testing"
)

func TestDivide(t *testing.T) {
	for _, c := range []struct {
		dividend int
		divisor  int
		expect   int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{1, 1, 1},
		{-2147483648, -1, 2147483647},
		{-2147483648, 1, -2147483648},
		{2147483647, 2, 2147483647 / 2},
	} {
		got := Divide(c.dividend, c.divisor)
		if got != c.expect {
			format := "dividend: %d\t|\tdivisor: %d\t|\texpect: %d\t|\tgot: %d"
			t.Errorf(format, c.dividend, c.divisor, c.expect, got)
		}
	}
}
