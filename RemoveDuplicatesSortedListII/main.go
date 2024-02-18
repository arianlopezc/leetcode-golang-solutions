package RemoveDuplicatesSortedListII

type ListNode struct {
	Val  int
	Next *ListNode
}

func lastNoDuplicated(node *ListNode) *ListNode {
	var nav = node.Next
	var swapped = false
	for {
		if nav != nil && node.Val == nav.Val {
			node = node.Next
			nav = node.Next
			swapped = true
		} else if nav == nil {
			if !swapped {
				return node
			} else {
				return nil
			}
		} else {
			if swapped {
				swapped = false
				node = node.Next
				nav = node.Next
			} else {
				return node
			}
		}
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := lastNoDuplicated(head)
	if first == nil {
		return nil
	}
	*head = *first
	nav := head
	for nav != nil && nav.Next != nil {
		nav.Next = lastNoDuplicated(nav.Next)
		nav = nav.Next
	}
	return head
}
