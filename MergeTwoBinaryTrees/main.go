package MergeTwoBinaryTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else {
		if root1 != nil && root2 == nil {
			return &TreeNode{Val: root1.Val, Left: mergeTrees(root1.Left, root1.Right), Right: nil}
		} else if root1 == nil && root2 != nil {
			return &TreeNode{Val: root1.Val, Right: mergeTrees(root2.Left, root2.Right), Left: nil}
		} else {
			return &TreeNode{Val: root1.Val + root2.Val, Left: mergeTrees(root1.Left, root1.Right), Right: mergeTrees(root2.Left, root2.Right)}
		}
	}
}
