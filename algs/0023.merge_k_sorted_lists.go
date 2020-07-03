package algs

/*
MergeKLists solves the following problem:
	Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

	Example:

		Input:
		[
		1->4->5,
		1->3->4,
		2->6
		]
		Output: 1->1->2->3->4->4->5->6
*/
func MergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	}

	first := MergeKLists(lists[:n/2])
	second := MergeKLists(lists[n/2:])
	return MergeTwoLists(first, second)
}
