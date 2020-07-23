package algs

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	for _, c := range []struct {
		nums   []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
	} {
		nums := make([]int, len(c.nums))
		copy(nums, c.nums)
		NextPermutation(nums)
		if !reflect.DeepEqual(nums, c.expect) {
			t.Errorf("nums: %v\t|\t expect: %v\t|\t got: %v", c.nums, c.expect, nums)
		}
	}
}
