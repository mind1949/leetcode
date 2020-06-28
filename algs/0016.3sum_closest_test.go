package algs

import "testing"

func TestThreeSumClosest(t *testing.T) {
	for _, c := range []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-5, -5, -4, -4, -3, 8}, 0, 0},
	} {
		got := ThreeSumClosest(c.nums, c.target)
		if c.expect != got {
			t.Errorf("nums: %v\t|\t target: %d \t|\texpect: %d\t|\tgot: %d", c.nums, c.target, c.expect, got)
		}
	}
}
