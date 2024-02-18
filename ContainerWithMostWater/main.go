package ContainerWithMostWater

func maxArea(height []int) int {
	var n = len(height)
	var l = 0
	var r = n - 1
	var ans = 0
	for l < r {
		var temp int
		if height[l] < height[r] {
			temp = height[l] * (r - l)
		} else {
			temp = height[r] * (r - l)
		}
		if temp > ans {
			ans = temp
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
