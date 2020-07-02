package algs

/*
MergeTwoLists solves the following problem:
	Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

	Example:

		Input: 1->2->4, 1->3->4
		Output: 1->1->2->3->4->4
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	guard := &ListNode{}
	head := guard
	for {
		if l1 == nil {
			head.Next = l2
			break
		}
		if l2 == nil {
			head.Next = l1
			break
		}

		if l1.Val < l2.Val {
			v := l1.Val
			for l1 != nil && v == l1.Val {
				node := l1
				l1 = l1.Next
				head.Next = node
				head = head.Next
			}
		} else {
			v := l2.Val
			for l2 != nil && v == l2.Val {
				node := l2
				l2 = l2.Next
				head.Next = node
				head = head.Next
			}
		}
	}

	return guard.Next
}
