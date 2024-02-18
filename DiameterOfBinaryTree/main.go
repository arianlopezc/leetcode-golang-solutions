package DiameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameter(largest *int, root *TreeNode) int {
	if root.Right == nil && root.Left == nil {
		return 1
	}
	var leftPath int
	if root.Left != nil {
		leftPath = diameter(largest, root.Left)
	}
	var rightPath int
	if root.Right != nil {
		rightPath = diameter(largest, root.Right)
	}
	var totalChildren = leftPath + rightPath
	if *largest < totalChildren {
		*largest = totalChildren
	}
	return max(leftPath, rightPath) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	var largest int
	diameter(&largest, root)
	return largest
}
