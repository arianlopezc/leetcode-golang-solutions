package RemoveDuplicatesSortedLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var nav = head
	for nav != nil {
		var skipper = nav
		for skipper.Next != nil && skipper.Next.Val == skipper.Val {
			skipper = skipper.Next
		}
		nav.Next = skipper.Next
		nav = nav.Next
	}
	return head
}
