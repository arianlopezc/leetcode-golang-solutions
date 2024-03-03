package MinimumDistanceBetweenBSTNodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var order = make([]int, 0)
	inOrder(root, &order)
	var minDiff = 10000000
	for index := 1; index < len(order); index++ {
		var diff = abs(order[index], order[index-1])
		if diff < minDiff {
			diff = minDiff
		}
	}
	return minDiff
}

func abs(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func inOrder(node *TreeNode, order *[]int) {
	if node != nil {
		inOrder(node.Left, order)
		*order = append(*order, node.Val)
		inOrder(node.Right, order)
	}
}
