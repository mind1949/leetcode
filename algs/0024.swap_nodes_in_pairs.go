package algs

/*
SwapPairs solves the following problem:
	Given a linked list, swap every two adjacent nodes and return its head.

	You may not modify the values in the list's nodes, only nodes itself may be changed.



	Example:

		Given 1->2->3->4, you should return the list as 2->1->4->3
*/
func SwapPairs(head *ListNode) *ListNode {
	guard := &ListNode{Next: head}
	for zero := guard; zero != nil && zero.Next != nil && zero.Next.Next != nil; zero = zero.Next.Next {
		first := zero.Next
		second := first.Next

		// swap
		first.Next = second.Next
		second.Next = first
		zero.Next = second
	}

	return guard.Next
}
