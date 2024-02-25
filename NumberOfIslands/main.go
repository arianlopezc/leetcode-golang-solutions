package NumberOfIslands

const (
	ONE byte = '1'
)

func numIslands(grid [][]byte) int {
	var islandId = 2
	var totalIslands int
	for indexRow := range grid {
		for indexCol := range grid[indexRow] {
			if grid[indexRow][indexCol] == ONE {
				tag(islandId, indexRow, indexCol, grid)
				islandId++
				totalIslands++
			}
		}
	}
	return totalIslands
}

func tag(island int, row int, col int, grid [][]byte) {
	grid[row][col] = byte(island)
	var lengthRow, lengthCol = len(grid), len(grid[0])
	if row-1 >= 0 {
		if row-1 >= 0 && grid[row-1][col] == ONE {
			tag(island, row-1, col, grid)
		}
	}
	if row+1 < lengthRow {
		if row+1 >= 0 && grid[row+1][col] == ONE {
			tag(island, row+1, col, grid)
		}
	}
	if col-1 >= 0 && grid[row][col-1] == ONE {
		tag(island, row, col-1, grid)
	}
	if col+1 < lengthCol && grid[row][col+1] == ONE {
		tag(island, row, col+1, grid)
	}
}
