package algs

// ListNode leetcode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
AddTwoNumbers solves following problem:

	You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

	You may assume the two numbers do not contain any leading zero, except the number 0 itself.

	Example:

		Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
		Output: 7 -> 0 -> 8
		Explanation: 342 + 465 = 807.
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	n1, n2, n3 := l1, l2, l3
	for n1 != nil || n2 != nil {
		if n1 != nil {
			n3.Val += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			n3.Val += n2.Val
			n2 = n2.Next
		}

		if n1 != nil || n2 != nil {
			n3.Next = &ListNode{}
		}
		if n3.Val >= 10 {
			if n3.Next == nil {
				n3.Next = &ListNode{}
			}
			n3.Next.Val = 1
			n3.Val -= 10
		}
		n3 = n3.Next
	}

	return l3
}

func ints2ListNode(ints []int) *ListNode {
	l := &ListNode{}
	n := l
	for i := 0; i < len(ints); i++ {
		n.Val = ints[i]
		if i < len(ints)-1 {
			n.Next = &ListNode{}
		}
		n = n.Next
	}
	return l
}

func listNode2Ints(head *ListNode) []int {
	ints := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		ints = append(ints, node.Val)
	}
	return ints
}
