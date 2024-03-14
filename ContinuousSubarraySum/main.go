package ContinuousSubarraySum

func checkSubarraySum(nums []int, k int) bool {
	mapped := make(map[int]int)
	mapped[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		if v, ok := mapped[sum]; ok {
			if i-v > 1 {
				return true
			}
		} else {
			mapped[sum] = i
		}
	}
	return false
}
