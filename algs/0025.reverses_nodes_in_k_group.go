package algs

/*
ReverseKGroup solves the following problem:

	Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

	k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

	Example:

		Given this linked list: 1->2->3->4->5

		For k = 2, you should return: 2->1->4->3->5

		For k = 3, you should return: 3->2->1->4->5

	Note:

		Only constant extra memory is allowed.
		You may not alter the values in the list's nodes, only nodes itself may be changed.
*/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	guard := &ListNode{Next: head}
	nodek := guard
	// 使nodek为第k个节点
	for i := 0; i < k; i++ {
		if nodek.Next == nil {
			return guard.Next
		}
		nodek = nodek.Next
	}

	zero := guard
	for {
		// 翻转
		first := zero.Next
		zero.Next = nodek
		l, r := first, first.Next // 相邻的两个节点（左右节点）
		first.Next = nodek.Next
		for i := 2; i <= k; i++ {
			tmp := r.Next
			r.Next = l // 交换左右两个节点

			l, r = r, tmp
		}

		// 位移k个节点
		zero = first
		nodek = first
		for i := 0; i < k; i++ {
			if nodek.Next == nil {
				return guard.Next
			}
			nodek = nodek.Next
		}
	}
}
