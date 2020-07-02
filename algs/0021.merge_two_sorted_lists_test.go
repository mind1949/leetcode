package algs

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	for _, c := range []struct {
		l1     *ListNode
		l2     *ListNode
		expect *ListNode
	}{
		{
			ints2ListNode([]int{1, 2, 4}),
			ints2ListNode([]int{1, 3, 4}),
			ints2ListNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			nil,
			ints2ListNode([]int{1, 3, 4}),
			ints2ListNode([]int{1, 3, 4}),
		},
		{
			ints2ListNode([]int{1, 2, 4, 6}),
			ints2ListNode([]int{1, 3, 4}),
			ints2ListNode([]int{1, 1, 2, 3, 4, 4, 6}),
		},
	} {
		got := MergeTwoLists(c.l1, c.l2)
		if !reflect.DeepEqual(
			listNode2Ints(got),
			listNode2Ints(c.expect),
		) {
			format := "l1: %v\t|\tl2: %v\t|\texpect: %v\t|\tgot: %v"
			t.Errorf(format, listNode2Ints(c.l1), listNode2Ints(c.l2), listNode2Ints(c.expect), listNode2Ints(got))
		}
	}
}
