package algs

import (
	"reflect"
	"testing"
)

/*
Input:

	[
	1->4->5,
	1->3->4,
	2->6
	]
	Output: 1->1->2->3->4->4->5->6
*/
func TestMergeKLists(t *testing.T) {
	for _, c := range []struct {
		input  []*ListNode
		expect *ListNode
	}{
		{
			[]*ListNode{ints2ListNode([]int{1, 4, 5}), ints2ListNode([]int{1, 3, 4}), ints2ListNode([]int{2, 6})},
			ints2ListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			[]*ListNode{ints2ListNode([]int{1, 4, 5})},
			ints2ListNode([]int{1, 4, 5}),
		},
	} {
		got := MergeKLists(c.input)
		if !reflect.DeepEqual(listNode2Ints(got), listNode2Ints(c.expect)) {
			t.Errorf("input: %v\t|\texpect: %v\t|\tgot: %v", c.input, listNode2Ints(c.expect), listNode2Ints(got))
		}
	}
}
