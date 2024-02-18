package IsValidBST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func validate(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	if node.Left != nil && node.Val <= node.Left.Val {
		return false
	}
	if node.Right != nil && node.Val >= node.Right.Val {
		return false
	}
	return true
}

func findMax(node *TreeNode, max *int) {
	if node != nil {
		if node.Val > *max {
			*max = node.Val
		}
		if node.Right != nil {
			findMax(node.Right, max)
		}
	}
}

func findMin(node *TreeNode, min *int) {
	if node != nil {
		if node.Val < *min {
			*min = node.Val
		}
		if node.Left != nil {
			findMin(node.Left, min)
		}
	}
}

func isValidBST(root *TreeNode) bool {
	if validate(root) {
		max := math.MinInt
		min := math.MaxInt
		findMax(root.Left, &max)
		findMin(root.Right, &min)
		if root.Val <= max {
			return false
		}
		if root.Val >= min {
			return false
		}
		var leftValid = true
		if root.Left != nil {
			leftValid = isValidBST(root.Left)
		}
		var rightValid = true
		if root.Right != nil {
			rightValid = isValidBST(root.Right)
		}
		return leftValid && rightValid
	}
	return false
}
