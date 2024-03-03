package MinCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	var short, long = cost[0], cost[1]
	for index := 2; index < len(cost); index++ {
		short, long = long, cost[index]+min(short, long)
	}
	return min(short, long)
}
