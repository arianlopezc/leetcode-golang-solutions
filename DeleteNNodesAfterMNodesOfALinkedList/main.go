package DeleteNNodesAfterMNodesOfALinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	var leftM = m - 1
	var mainNav = head
	for mainNav != nil {
		if leftM == 0 {
			var secondNav = mainNav
			var leftN = n
			for leftN >= 0 && secondNav != nil {
				secondNav = secondNav.Next
				leftN--
			}
			mainNav.Next = secondNav
			leftM = m
		}
		mainNav = mainNav.Next
		leftM--
	}
	return head
}
