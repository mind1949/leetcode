package algs

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	for _, c := range []struct {
		nums    []int
		expect1 int
		expect2 []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{1, 2, 3, 3, 4, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
	} {
		got := RemoveDuplicates(c.nums)
		if got != c.expect1 {
			t.Errorf("nums: %v\t|\texpect: %d\t|\tgot: %d", c.nums, c.expect1, got)
		}
		if !reflect.DeepEqual(c.nums[:got], c.expect2) {
			t.Errorf("nums: %v\t|\texpect: %v\t|\tgot: %v", c.nums, c.expect2, c.nums[:got])
		}
	}
}
