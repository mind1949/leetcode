package algs

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for _, c := range []struct {
		nums    []int
		val     int
		expect1 int
		expect2 []int
	}{
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
		{nil, 1, 0, nil},
	} {
		got := RemoveElement(c.nums, c.val)
		if got != c.expect1 {
			t.Errorf("nums: %v\t|\tval: %d\t|\texpect %d\t|\tgot: %d", c.nums, c.val, c.expect1, got)
		}
		if reflect.DeepEqual(c.nums[:got], c.expect2) {

		}
	}
}
