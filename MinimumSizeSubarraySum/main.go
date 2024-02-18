package MinimumSizeSubarraySum

import "math"

func minSubArrayLen(target int, nums []int) int {
	var shortest = math.MaxInt32
	var currentShortest, front, rear, ongoingSum int
	var length = len(nums)
	for {
		if ongoingSum >= target {
			if currentShortest < shortest {
				shortest = currentShortest
			}
			ongoingSum -= nums[front]
			currentShortest--
			front++
		} else if ongoingSum < target && front <= rear && rear < length {
			ongoingSum += nums[rear]
			currentShortest++
			rear++
		} else {
			if shortest == math.MaxInt32 {
				return 0
			}
			return shortest
		}
	}
}
