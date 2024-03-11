package ShortestPathInBinaryMatrix

type Point struct {
	Row int
	Col int
}

func getChildren(point Point, grid [][]int) []Point {
	var children []Point
	if point.Row-1 >= 0 {
		if grid[point.Row-1][point.Col] == 0 {
			children = append(children, Point{
				Row: point.Row - 1,
				Col: point.Col,
			})
		}
		if point.Col-1 >= 0 && grid[point.Row-1][point.Col-1] == 0 {
			children = append(children, Point{
				Row: point.Row - 1,
				Col: point.Col - 1,
			})
		}
		if point.Col+1 < len(grid[point.Row-1]) && grid[point.Row-1][point.Col+1] == 0 {
			children = append(children, Point{
				Row: point.Row - 1,
				Col: point.Col + 1,
			})
		}
	}
	if point.Row+1 < len(grid) {
		if grid[point.Row+1][point.Col] == 0 {
			children = append(children, Point{
				Row: point.Row + 1,
				Col: point.Col,
			})
		}
		if point.Col-1 >= 0 && grid[point.Row+1][point.Col-1] == 0 {
			children = append(children, Point{
				Row: point.Row + 1,
				Col: point.Col - 1,
			})
		}
		if point.Col+1 < len(grid[point.Row+1]) && grid[point.Row+1][point.Col+1] == 0 {
			children = append(children, Point{
				Row: point.Row + 1,
				Col: point.Col + 1,
			})
		}
	}
	if len(grid[point.Row]) > point.Col+1 && grid[point.Row][point.Col+1] == 0 {
		children = append(children, Point{
			Row: point.Row,
			Col: point.Col + 1,
		})
	}
	if 0 <= point.Col-1 && grid[point.Row][point.Col-1] == 0 {
		children = append(children, Point{
			Row: point.Row,
			Col: point.Col - 1,
		})
	}
	return children
}

func shortestPathBinaryMatrix(grid [][]int) int {
	var mapped = make(map[Point]int)
	var queue = make([]Point, 0)
	if grid[0][0] == 0 {
		var startingPoint = Point{
			Row: 0,
			Col: 0,
		}
		queue = append(queue, startingPoint)
		mapped[startingPoint] = 1
	}
	for len(queue) > 0 {
		var point = queue[0]
		queue = queue[1:]
		var children = getChildren(point, grid)
		var pointTotal = mapped[point] + 1
		for _, child := range children {
			if val, ok := mapped[child]; !ok {
				mapped[child] = pointTotal
				queue = append(queue, child)
			} else {
				if val >= pointTotal {
					mapped[child] = pointTotal
					queue = append(queue, child)
				}
			}
		}
	}
	if val, ok := mapped[Point{
		Row: len(grid) - 1,
		Col: len(grid[len(grid)-1]) - 1,
	}]; ok {
		return val
	} else {
		return -1
	}
}
