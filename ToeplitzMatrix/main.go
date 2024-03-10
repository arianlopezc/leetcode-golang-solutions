package ToeplitzMatrix

func isToeplitzMatrix(matrix [][]int) bool {
	for indexRow := len(matrix) - 2; indexRow >= 0; indexRow-- {
		var pos int
		var val = matrix[indexRow][pos]
		for indexCol := pos; indexRow+pos < len(matrix) && indexCol+pos < len(matrix[indexRow]); pos++ {
			if val != matrix[indexRow+pos][indexCol+pos] {
				return false
			}
		}
	}
	for indexCol := 1; indexCol < len(matrix[0]); indexCol++ {
		var pos int
		var val = matrix[pos][indexCol]
		for indexRow := pos; indexRow+pos < len(matrix) && indexCol+pos < len(matrix[indexRow]); pos++ {
			if val != matrix[indexRow+pos][indexCol+pos] {
				return false
			}
		}
	}
	return true
}
