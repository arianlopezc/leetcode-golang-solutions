package AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var headResult *ListNode
	var nextResult = &ListNode{}
	headResult = nextResult
	for l1 != nil || l2 != nil {
		var addedVal int
		if l1 != nil {
			addedVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			addedVal += l2.Val
			l2 = l2.Next
		}
		addedVal += carry
		if addedVal > 9 {
			carry = 1
			addedVal = addedVal % 10
		} else {
			carry = 0
		}
		newNode := &ListNode{Val: addedVal}
		nextResult.Next = newNode
		nextResult = nextResult.Next
	}
	if carry > 0 {
		newNode := &ListNode{Val: carry}
		nextResult.Next = newNode
		nextResult = nextResult.Next
	}
	return headResult.Next
}
