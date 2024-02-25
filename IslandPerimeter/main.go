package IslandPerimeter

func islandPerimeter(grid [][]int) int {
	var totalPerimeter int
	for rowIndex := range grid {
		for colIndex := range grid[rowIndex] {
			if grid[rowIndex][colIndex] == 1 {
				totalPerimeter += totalPerimeterInPoint(rowIndex, colIndex, grid)
			}
		}
	}
	return totalPerimeter
}

func totalPerimeterInPoint(x int, y int, grid [][]int) int {
	var lengthRow, lengthCol, total = len(grid), len(grid[0]), 0
	if x+1 == lengthRow {
		total++
	}
	if x-1 == -1 {
		total++
	}
	if x+1 < lengthRow && grid[x+1][y] == 0 {
		total++
	}
	if x-1 >= 0 && grid[x-1][y] == 0 {
		total++
	}
	if y+1 == lengthCol {
		total++
	}
	if y-1 == -1 {
		total++
	}
	if y+1 < lengthCol && grid[x][y+1] == 0 {
		total++
	}
	if y-1 >= 0 && grid[x][y-1] == 0 {
		total++
	}
	return total
}
