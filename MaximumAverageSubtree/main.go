package MaximumAverageSubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageSubtree(node *TreeNode, largest *float64) (int, int) {
	if node != nil {
		var totalLeft, quantLeft = averageSubtree(node.Left, largest)
		var totalRight, quantRight = averageSubtree(node.Right, largest)
		if quantLeft+quantRight == 0 {
			var self = float64(node.Val)
			if self > *largest {
				*largest = self
			}
		}
		var combined = float64(totalLeft+totalRight+node.Val) / float64(quantLeft+quantRight+1)
		if combined > *largest {
			*largest = combined
		}
		return totalLeft + totalRight + node.Val, quantLeft + quantRight + 1
	}
	return 0.0, 0
}

func maximumAverageSubtree(root *TreeNode) float64 {
	var largest float64
	averageSubtree(root, &largest)
	return largest
}
