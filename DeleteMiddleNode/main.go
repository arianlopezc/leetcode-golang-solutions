package DeleteMiddleNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head != nil && head.Next == nil {
		return nil
	}
	if head != nil && head.Next != nil && head.Next.Next == nil {
		head.Next = nil
		return head
	}
	var nav *ListNode = head
	var aheadNav *ListNode = nav.Next.Next
	for aheadNav != nil {
		if aheadNav.Next != nil && aheadNav.Next.Next != nil {
			aheadNav = aheadNav.Next.Next
			nav = nav.Next
		} else if aheadNav.Next != nil && aheadNav.Next.Next == nil {
			aheadNav = nil
			nav = nav.Next
		} else {
			aheadNav = nil
		}
	}
	*nav.Next = *nav.Next.Next
	return head
}
