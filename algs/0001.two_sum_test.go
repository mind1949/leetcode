package algs

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	cs := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 5, 1, 6, 7},
			6,
			[]int{1, 2},
		},
	}

	for _, c := range cs {
		got := TwoSum(c.nums, c.target)
		if got[0] != c.want[0] || got[1] != c.want[1] {
			format := "nums: %v | target: %d | want: %v | got: %v"
			t.Errorf(format, c.nums, c.target, c.want, got)
		}
	}
}
