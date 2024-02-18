package HasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var front, rear = head, head.Next.Next
	for front != nil {
		if front.Val == rear.Val {
			return true
		} else {
			rear = rear.Next
			if front.Next != nil {
				front = front.Next.Next
			} else {
				front = front.Next
			}
		}
	}
	return false
}
