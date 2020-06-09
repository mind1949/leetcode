package algs

import (
	"reflect"
	"testing"
)

func TestInts2ListNode(t *testing.T) {
	cs := []struct {
		ints []int
		l    *ListNode
	}{
		{
			ints: []int{5, 7, 8},
			l: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	for i, c := range cs {
		l := intSlice2ListNode(c.ints)
		if !reflect.DeepEqual(l, c.l) {
			t.Errorf("testCase: %d", i)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	cs := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   intSlice2ListNode([]int{5, 7, 8}),
			l2:   intSlice2ListNode([]int{3, 2, 7, 9}),
			want: intSlice2ListNode([]int{8, 9, 5, 0, 1}),
		},
		{
			l1:   intSlice2ListNode([]int{1, 8}),
			l2:   intSlice2ListNode([]int{0}),
			want: intSlice2ListNode([]int{1, 8}),
		},
	}

	for i, c := range cs {
		got := AddTwoNumbers(c.l1, c.l2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("testCase: %d", i)
		}
	}
}
