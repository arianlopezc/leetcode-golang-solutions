package MiddleOfTheLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var nav, navNext *ListNode
	nav = head
	navNext = nav
	for navNext != nil {
		nav = nav.Next
		if navNext.Next == nil || navNext.Next.Next == nil {
			navNext = nil
		} else {
			navNext = navNext.Next.Next
			if navNext.Next == nil {
				return nav
			}
		}
	}
	return nav
}
