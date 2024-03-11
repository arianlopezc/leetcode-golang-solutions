package KthSmallestElementInASortedMatrix

func kthSmallest(matrix [][]int, k int) int {
	var merged []int
	var row int
	for len(matrix) != row {
		var indexA, indexB int
		var temp = make([]int, len(merged)+len(matrix[row]))
		var colIndex int
		for indexA < len(merged) || indexB < len(matrix[row]) {
			if indexA >= len(merged) {
				temp[colIndex] = matrix[row][indexB]
				indexB++
			} else if indexB >= len(matrix[row]) {
				temp[colIndex] = merged[indexA]
				indexA++
			} else {
				if merged[indexA] <= matrix[row][indexB] {
					temp[colIndex] = merged[indexA]
					indexA++
				} else {
					temp[colIndex] = matrix[row][indexB]
					indexB++
				}
			}
			colIndex++
		}
		merged = temp
		row++
	}
	return merged[k-1]
}
