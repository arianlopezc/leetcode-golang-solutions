package BinaryTreePruning

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func prune(node *TreeNode) (*TreeNode, bool) {
	if node != nil {
		var _, leftDo = prune(node.Left)
		var _, rightDo = prune(node.Right)
		if !leftDo {
			node.Left = nil
		}
		if !rightDo {
			node.Right = nil
		}
		if node.Val == 1 {
			return node, true
		}
		if leftDo || rightDo {
			return node, true
		} else {
			return nil, false
		}
	}
	return node, false
}

func pruneTree(root *TreeNode) *TreeNode {
	var node, _ = prune(root)
	return node
}
