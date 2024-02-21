package IsPalindromeLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func palindrome(head *ListNode, node *ListNode) (isPal bool, sameNode *ListNode) {
	if node.Next == nil {
		return node.Val == head.Val, head.Next
	} else {
		isPal, sameNode := palindrome(head, node.Next)
		if !isPal {
			return false, sameNode.Next
		} else {
			return node.Val == sameNode.Val, sameNode.Next
		}
	}
}

func isPalindrome(head *ListNode) bool {
	isPal, _ := palindrome(head, head)
	return isPal
}
