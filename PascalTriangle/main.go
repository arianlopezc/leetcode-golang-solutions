package PascalTriangle

func generate(numRows int) [][]int {
	var result = [][]int{{1}}
	numRows--
	var rowTotal = 1
	for numRows > 0 {
		result = append(result, make([]int, rowTotal+1))
		for index := 0; index <= rowTotal; index++ {
			var left int
			if index-1 >= 0 {
				left = result[rowTotal-1][index-1]
			}
			var right int
			if index < rowTotal {
				right = result[rowTotal-1][index]
			}
			result[rowTotal][index] = left + right
		}
		rowTotal++
		numRows--
	}
	return result
}
