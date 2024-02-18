package LongestSubarrayOfOnes

func longestSubarray(nums []int) int {
	var maxOnes, totalInWindow = 0, 0
	var zeroUsed = false
	for front, rear := 0, 0; front < len(nums); {
		if nums[front] == 1 || !zeroUsed {
			if nums[front] == 0 {
				zeroUsed = true
			} else {
				totalInWindow++
				if totalInWindow > maxOnes {
					maxOnes = totalInWindow
				}
			}
			front++
		} else {
			if nums[rear] == 0 {
				zeroUsed = false
			} else {
				totalInWindow--
			}
			rear++
		}
	}
	if maxOnes == len(nums) {
		maxOnes--
	}
	return maxOnes
}
