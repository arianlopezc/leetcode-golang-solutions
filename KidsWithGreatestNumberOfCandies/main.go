package KidsWithGreatestNumberOfCandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return []bool{}
	}
	var max int = -1
	for _, value := range candies {
		if max < value {
			max = value
		}
	}
	var result []bool = make([]bool, len(candies))
	for index, value := range candies {
		if value+extraCandies >= max {
			result[index] = true
		}
	}
	return result
}
