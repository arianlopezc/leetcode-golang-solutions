package UniqueNumberOfOccurrences

func uniqueOccurrences(arr []int) bool {
	var instances map[int]int = make(map[int]int)
	for _, value := range arr {
		instances[value]++
	}
	var set map[int]bool = make(map[int]bool)
	for _, data := range instances {
		if set[data] {
			return false
		} else {
			set[data] = true
		}
	}
	return true
}
