package IntersectionOfTwoLinkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var mapped = make(map[*ListNode]bool)
	for headA != nil {
		mapped[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := mapped[headB]; ok {
			return headB
		} else {
			mapped[headB] = true
			headB = headB.Next
		}
	}
	return nil
}
