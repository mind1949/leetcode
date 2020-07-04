package algs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	for _, c := range []struct {
		input  *ListNode
		expect *ListNode
	}{
		{ints2ListNode([]int{1, 2, 3, 4}), ints2ListNode([]int{2, 1, 4, 3})},
		{nil, nil},
		{ints2ListNode([]int{1}), ints2ListNode([]int{1})},
	} {
		got := SwapPairs(c.input)
		if !reflect.DeepEqual(listNode2Ints(got), listNode2Ints(c.expect)) {
			t.Errorf("input: %v\r|\rexpect: %v\r|\rgot: %v\n", listNode2Ints(c.input), listNode2Ints(c.expect), listNode2Ints(got))
		}
	}
}
