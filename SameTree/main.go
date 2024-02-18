package SameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val != q.Val || p != nil && q == nil || p == nil && q != nil || p.Left != nil && q.Left == nil || p.Right != nil && q.Right == nil || p.Left == nil && q.Left != nil || p.Right == nil && q.Right != nil {
		return false
	}
	if p.Left != nil {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
	}
	if p.Right != nil {
		if !isSameTree(p.Right, q.Right) {
			return false
		}
	}
	return true
}
