package algs

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	for _, c := range []struct {
		nums   []int
		target int
		expect [][]int
	}{
		/*
			[
			  [-1,  0, 0, 1],
			  [-2, -1, 1, 2],
			  [-2,  0, 0, 2]
			]
		*/
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			[]int{0, 0, 0, 0},
			0,
			[][]int{
				{0, 0, 0, 0},
			},
		},
	} {
		got := FourSum(c.nums, c.target)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("nums: %v\t|\ttarget: %d\t|\texpect: %v\t|\tgot: %v", c.nums, c.target, c.expect, got)
		}
	}
}
