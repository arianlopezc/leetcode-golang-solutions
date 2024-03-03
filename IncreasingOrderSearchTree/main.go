package IncreasingOrderSearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func transform(root *TreeNode, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	res := transform(root.Left, root)
	root.Left = nil
	root.Right = transform(root.Right, tail)
	return res
}

func increasingBST(root *TreeNode) *TreeNode {
	return transform(root, nil)
}
