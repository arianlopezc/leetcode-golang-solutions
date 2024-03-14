package LowestCommonAncestorOfABinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val || (root.Val > min(p.Val, q.Val) && root.Val < max(p.Val, q.Val)) {
		return root
	} else {
		if root.Val > p.Val {
			return lowestCommonAncestor(root.Left, p, q)
		} else {
			return lowestCommonAncestor(root.Right, p, q)
		}
	}
}
