package OddEvenLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	var startOdd ListNode = ListNode{}
	var navOdd *ListNode = &startOdd
	var nav *ListNode = head
	for nav != nil && nav.Next != nil {
		navOdd.Next = nav.Next
		nav.Next = nav.Next.Next
		navOdd = navOdd.Next
		navOdd.Next = nil
		if nav.Next == nil {
			nav.Next = startOdd.Next
			nav = nil
		} else {
			nav = nav.Next
		}
	}
	if nav != nil {
		nav.Next = startOdd.Next
	}
	return head
}
