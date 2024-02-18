package SortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(navOne *ListNode, navTwo *ListNode) *ListNode {
	var start = &ListNode{}
	var nav = start
	for navOne != nil || navTwo != nil {
		if navOne == nil {
			nav.Next = navTwo
			navTwo = navTwo.Next
		} else if navTwo == nil {
			nav.Next = navOne
			navOne = navOne.Next
		} else {
			if navOne.Val <= navTwo.Val {
				nav.Next = navOne
				navOne = navOne.Next
			} else {
				nav.Next = navTwo
				navTwo = navTwo.Next
			}
		}
		nav = nav.Next
	}
	return start.Next
}

func splitAll(nav *ListNode) []*ListNode {
	var splitted = make([]*ListNode, 0)
	for nav != nil {
		var temp *ListNode
		temp, nav = nav, nav.Next
		temp.Next = nil
		splitted = append(splitted, temp)
	}
	return splitted
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var divided = splitAll(head)
	for len(divided) > 1 {
		var first = divided[0]
		var second = divided[1]
		divided = append(divided[2:], merge(first, second))
	}
	return divided[0]
}
