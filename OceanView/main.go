package OceanView

import "slices"

func findBuildings(heights []int) []int {
	var withView = []int{len(heights) - 1}
	var currTallest = heights[len(heights)-1]
	var navIndex = len(heights) - 1
	for navIndex >= 0 {
		if heights[navIndex] > currTallest {
			currTallest = heights[navIndex]
			withView = append(withView, navIndex)
		}
		navIndex--
	}
	slices.Sort(withView)
	return withView
}
