package MaxTwinSumLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var nums []int
	var max int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	var length = len(nums)
	for i, val := range nums {
		if sum := val + nums[length-i-1]; max < sum {
			max = sum
		}
	}
	return max
}
