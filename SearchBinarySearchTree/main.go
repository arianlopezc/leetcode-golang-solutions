package SearchBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if val < root.Val {
			return searchBST(root.Left, val)
		}
		if val > root.Val {
			return searchBST(root.Right, val)
		}
	}
	return nil
}
