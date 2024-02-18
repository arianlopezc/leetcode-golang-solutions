package RemoveLinkedListElements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var navIndex = head
	var result = head
	for navIndex != nil {
		if navIndex.Next != nil && navIndex.Next.Val == val {
			navIndex.Next = navIndex.Next.Next
		} else {
			navIndex = navIndex.Next
		}
	}
	if result != nil && result.Val == val {
		return result.Next
	}
	return result
}
