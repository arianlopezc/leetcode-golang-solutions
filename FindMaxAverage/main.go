package FindMaxAverage

func findMaxAverage(nums []int, k int) float64 {
	if k == 0 {
		return 0
	}
	var toAdd float64 = float64(k)
	var rearIndex int = 0
	var frontIndex int = 0
	var biggestAverage float64 = -10000000
	var currTotal float64 = 0
	var length int = len(nums)
	for frontIndex <= length {
		if k > 0 {
			if frontIndex < length {
				currTotal += float64(nums[frontIndex])
				k--
			}
			frontIndex++
		} else {
			newAve := currTotal / toAdd
			if newAve > biggestAverage {
				biggestAverage = newAve
			}
			currTotal -= float64(nums[rearIndex])
			rearIndex++
			k++
		}
	}
	return biggestAverage
}
