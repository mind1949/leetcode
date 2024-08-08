package algs

// IsPalindromeList
//
// Given the head of a singly linked list, return true if it is a
// palindrome
// or false otherwise.
//
// Example 1:
//
// Input: head = [1,2,2,1]
// Output: true
//
// Example 2:
//
// Input: head = [1,2]
// Output: false
//
// Constraints:
//
//	The number of nodes in the list is in the range [1, 10^5].
//	0 <= Node.val <= 9
//
// Follow up: Could you do it in O(n) time and O(1) space?
func IsPalindromeList(head *ListNode) bool {
	// 求长度
	length := 0
	for next := head; next != nil; next = next.Next {
		length++
	}

	// 找到后半部分开头
	cur := head
	for i := 0; i < length/2; i++ {
		cur = cur.Next
	}
	// 反转链表后半部分
	next := cur.Next
	for next != nil {
		tmp := next.Next
		next.Next = cur

		cur = next
		next = tmp
	}

	// 遍历头尾互相比较
	left, right := head, cur
	for i := 0; i < length/2; i++ {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}
	return true
}
