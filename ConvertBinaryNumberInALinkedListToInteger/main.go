package ConvertBinaryNumberInALinkedListToInteger

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var binRep string
	for head != nil {
		binRep += strconv.Itoa(head.Val)
		head = head.Next
	}
	var res, _ = strconv.ParseInt(binRep, 2, 64)
	return int(res)
}
