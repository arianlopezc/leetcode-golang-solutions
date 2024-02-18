package ClosestBinarySearchTreeValue

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closest(root *TreeNode, target *float64, smallest *int) {
	if float64(root.Val) == *target {
		*smallest = root.Val
	} else {
		var diff = math.Abs(*target - float64(root.Val))
		var smallestDiff = math.Abs(*target - float64(*smallest))
		if diff == smallestDiff {
			*smallest = min(root.Val, *smallest)
		} else if diff < smallestDiff {
			*smallest = root.Val
		}
		if root.Left != nil && float64(root.Val) > *target {
			closest(root.Left, target, smallest)
		}
		if root.Right != nil && float64(root.Val) < *target {
			closest(root.Right, target, smallest)
		}
	}
}

func closestValue(root *TreeNode, target float64) int {
	var smallest = math.MaxInt
	closest(root, &target, &smallest)
	return smallest
}
