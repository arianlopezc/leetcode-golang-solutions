package ProductOfArrayExceptSelf

func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	for index := range result {
		result[index] = 1
	}
	first, third := 1, 1
	indexFront, indexRear := 0, length-1
	for indexFront < length {
		result[indexFront] = first * result[indexFront]
		first = nums[indexFront] * result[indexFront]
		indexFront++
	}
	for indexRear >= 0 {
		result[indexRear] = third * result[indexRear]
		third = nums[indexRear] * third
		indexRear--
	}
	return result
}
