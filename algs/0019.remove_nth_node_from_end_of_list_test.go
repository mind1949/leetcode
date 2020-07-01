package algs

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	for _, c := range []struct {
		head   *ListNode
		n      int
		expect *ListNode
	}{
		{ints2ListNode([]int{1, 2, 3, 4, 5}), 2, ints2ListNode([]int{1, 2, 3, 5})},
		{ints2ListNode([]int{2, 1}), 2, ints2ListNode([]int{1})},
	} {
		got := removeNthFromEnd(c.head, c.n)
		if !reflect.DeepEqual(listNode2Ints(got), listNode2Ints(c.expect)) {
			format := "head: %v\t|\tn: %d\t|\texpect: %v\t|\tgot: %v"
			t.Errorf(format, listNode2Ints(c.head), c.n, listNode2Ints(c.expect), listNode2Ints(got))
		}
	}
}
