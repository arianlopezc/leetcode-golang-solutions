package HubSpot

func exercise(a []int, b []int, maxLength int) []int {
	var length int = maxLength
	if len(a)+len(b) < maxLength {
		length = len(a) + len(b)
	}
	var result []int = make([]int, length)
	var indexA int
	var indexB int
	for index, _ := range result {
		if indexA == len(a) {
			result[index] = b[indexB]
			indexB++
		} else if indexB == len(b) {
			result[index] = a[indexA]
			indexA++
		} else if a[indexA] <= b[indexB] {
			result[index] = a[indexA]
			indexA++
		} else {
			result[index] = b[indexB]
			indexB++
		}
	}
	return result
}
