package RemoveElement

func removeElement(nums []int, val int) int {
	for index := 0; index < len(nums); index++ {
		if nums[index] == val {
			nums = append(nums[:index], nums[index+1:]...)
			index--
		}
	}
	return len(nums)
}
