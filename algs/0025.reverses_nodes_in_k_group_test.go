package algs

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	for _, c := range []struct {
		head   *ListNode
		k      int
		expect *ListNode
	}{
		{
			head:   ints2ListNode([]int{1, 2, 3, 4, 5}),
			k:      6,
			expect: ints2ListNode([]int{1, 2, 3, 4, 5}),
		},
		{
			head:   ints2ListNode([]int{1, 2, 3, 4, 5}),
			k:      2,
			expect: ints2ListNode([]int{2, 1, 4, 3, 5}),
		},
		{
			head:   ints2ListNode([]int{1, 2, 3, 4, 5}),
			k:      3,
			expect: ints2ListNode([]int{3, 2, 1, 4, 5}),
		},
	} {
		got := ReverseKGroup(c.head, c.k)
		if !reflect.DeepEqual(listNode2Ints(c.expect), listNode2Ints(got)) {
			format := "head: %v | k: %d | expect: %v | got: %v"
			t.Errorf(format, listNode2Ints(c.head), c.k, listNode2Ints(c.expect), listNode2Ints(got))
		}
	}
}
