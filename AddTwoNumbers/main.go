package AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func invert(list *ListNode) *ListNode {
	var prev *ListNode
	var current = list
	for current != nil {
		var next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = invert(l1)
	l2 = invert(l2)
	var nav *ListNode
	var carryOver int
	for l1 != nil || l2 != nil {
		var valL1 int
		var valL2 int
		if l1 != nil {
			valL1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			valL2 = l2.Val
			l2 = l2.Next
		}
		var added = valL1 + valL2 + carryOver
		if added >= 10 {
			added = added % 10
			carryOver = 1
		} else {
			carryOver = 0
		}
		var newOne = ListNode{Val: added, Next: nav}
		nav = &newOne
	}
	if carryOver > 0 {
		var newOne = ListNode{Val: carryOver, Next: nav}
		nav = &newOne
	}
	return invert(nav)
}

//func queryNumberFromSlice(node *ListNode) int {
//	var decimalPoint = 1
//	var number int
//	for node != nil {
//		number = number + node.Val*decimalPoint
//		node = node.Next
//		decimalPoint *= 10
//	}
//	return number
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var added = queryNumberFromSlice(l1) + queryNumberFromSlice(l2)
//	var nav *ListNode
//	for _, val := range strings.Split(strconv.Itoa(added), "") {
//		if v, err := strconv.Atoi(val); err == nil {
//			var newOne = ListNode{Val: v, Next: nav}
//			nav = &newOne
//		}
//	}
//	return nav
//}
