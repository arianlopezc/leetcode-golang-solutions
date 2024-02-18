package FindPivotIndex

func pivotIndex(nums []int) int {
	var totalRight int
	for _, value := range nums {
		totalRight += value
	}
	var totalLeft int
	for index, value := range nums {
		if totalRight-value == totalLeft {
			return index
		} else {
			totalLeft += value
			totalRight -= value
		}
	}
	return -1
}
