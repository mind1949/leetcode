package algs

import (
	"reflect"
	"testing"
)

func TestTreeSum(t *testing.T) {
	for _, c := range []struct {
		input  []int
		expect [][]int
	}{
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			[]int{0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
	} {
		got := ThreeSum(c.input)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("input: %v\t|\texpect: %v\t|\tgot: %v", c.input, c.expect, got)
		}
	}
}
