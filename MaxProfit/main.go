package MaxProfit

func maxProfit(prices []int) int {
	var maxProfit = 0
	var rear, front, length = 0, 1, len(prices)
	for front < length {
		if prices[rear] < prices[front] {
			profit := prices[front] - prices[rear]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			rear = front
		}
		front++
	}
	return maxProfit
}
