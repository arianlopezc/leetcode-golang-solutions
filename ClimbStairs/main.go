package ClimbStairs

func climbStairs(n int) int {
	var firstWay = 0
	var secondWay = 1
	var tempWay = 0
	for i := 0; i < n; i++ {
		tempWay = firstWay + secondWay
		firstWay = secondWay
		secondWay = tempWay
	}
	return secondWay
}
