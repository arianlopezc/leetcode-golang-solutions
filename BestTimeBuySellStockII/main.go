package BestTimeBuySellStockII

func maxProfit(prices []int) int {
	var total int
	var index int
	for index < len(prices)-1 {
		if prices[index] < prices[index+1] {
			total += prices[index+1] - prices[index]
		}
	}
	return total
}
