package RemoveNthElementFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

func remove(node *ListNode, parent *ListNode, n int) (*ListNode, int) {
	if node == nil {
		return nil, 0
	}
	var myPos int
	var next *ListNode
	if node.Next != nil {
		next, myPos = remove(node.Next, node, n)
	}
	myPos++
	if n == myPos {
		if parent == nil {
			return next, myPos
		}
		parent.Next = next
		return parent.Next, myPos
	}
	return node, myPos
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var res, _ = remove(head, nil, n)
	return res
}
