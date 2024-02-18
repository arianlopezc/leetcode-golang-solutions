package IsSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func areSymmetric(nodeOne *TreeNode, nodeTwo *TreeNode) bool {
	if nodeOne == nil && nodeTwo == nil {
		return true
	}
	if nodeOne == nil && nodeTwo != nil || nodeOne != nil && nodeTwo == nil || nodeOne.Val != nodeTwo.Val {
		return false
	}
	if nodeOne.Left == nil && nodeTwo.Right != nil || nodeOne.Right == nil && nodeTwo.Left != nil {
		return false
	}
	return areSymmetric(nodeOne.Left, nodeTwo.Right) && areSymmetric(nodeOne.Right, nodeTwo.Left)
}

func isSymmetric(root *TreeNode) bool {
	return areSymmetric(root.Left, root.Right)
}
