package TwoSum

func twoSum(nums []int, target int) []int {
	mapped := make(map[int]int)
	for index, val := range nums {
		diff := target - val
		if diffVal, ok := mapped[diff]; ok {
			return []int{diffVal, index}
		} else {
			mapped[val] = index
		}
	}
	return []int{-1, -1}
}
