package algs

import "testing"

func TestMaxArea(t *testing.T) {
	cs := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, c := range cs {
		got := MaxArea(c.input)
		if got != c.expect {
			t.Errorf("input: %v \t|\texpect: %d\t|\tgot: %d", c.input, c.expect, got)
		}
	}
}
