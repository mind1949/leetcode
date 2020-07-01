package algs

/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 求出链表的长度
	listLen := 1
	for cur := head; cur.Next != nil; cur = cur.Next {
		listLen++
	}
	// 应对链表：[2, 1]，n=2的情况
	if listLen == n {
		return head.Next
	}
	// 计算倒数n+1与n-1个节点
	var (
		nplus  = &ListNode{} // 倒数第n+1个节点
		nminus = &ListNode{} // 倒数第n-1个节点
	)
	for cur, i := head, 0; i <= listLen; cur, i = cur.Next, i+1 {
		if i == listLen-n-1 {
			nplus = cur
		}
		if i == listLen-n+1 {
			nminus = cur
			break
		}
	}
	// 删除倒数第n个节点
	nplus.Next = nminus

	return head
}
