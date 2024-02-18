package MergeTwoSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var curr, dummy *ListNode
	dummy = &ListNode{}
	curr = dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1, curr = list1.Next, list1
		} else {
			curr.Next = list2
			list2, curr = list2.Next, list2
		}
	}
	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}
	return dummy.Next
}
