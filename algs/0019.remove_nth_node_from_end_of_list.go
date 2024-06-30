package algs

/*
RemoveNthFromEnd solves the following problem:

	Given a linked list, remove the n-th node from the end of list and return its head.

	Example:

		Given linked list: 1->2->3->4->5, and n = 2.

		After removing the second node from the end, the linked list becomes 1->2->3->5.

	Note:

	Given n will always be valid.

	Follow up:

	Could you do this in one pass?
*/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	guard := &ListNode{Next: head}

	first := guard
	second := guard
	for i := 0; i < n; i++ {
		second = second.Next
	}
	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next

	return guard.Next
}
