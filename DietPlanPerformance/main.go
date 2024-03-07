package DietPlanPerformance

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	var front, rear, total, ongoing int
	for front < len(calories) {
		for k > 0 {
			ongoing += calories[front]
			k--
			front++
		}
		if ongoing < lower {
			total--
		} else if ongoing > upper {
			total++
		}
		k++
		ongoing -= calories[rear]
		rear++
	}
	return total
}
