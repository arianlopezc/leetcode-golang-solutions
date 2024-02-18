package KthMissingPositiveNumber

func findKthPositive(arr []int, k int) int {
	var left, right = 0, len(arr) - 1
	var middle int
	for left <= right {
		middle = left + (right-left)/2
		var missingAtThisPoint = arr[middle] - middle - 1
		if missingAtThisPoint < k {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	var missingAtThisPoint = arr[middle] - middle - 1
	if missingAtThisPoint >= k {
		return arr[middle] - ((missingAtThisPoint - k) + 1)
	} else {
		return arr[middle] + (k - missingAtThisPoint)
	}
}
