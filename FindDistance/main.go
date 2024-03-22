package FindDistance

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceTo(node *TreeNode, val int) int {
	if node != nil {
		if node.Val == val {
			return 0
		} else {
			var left = distanceTo(node.Left, val)
			if left != -1 {
				return left + 1
			}
			var right = distanceTo(node.Right, val)
			if right != -1 {
				return right + 1
			}
		}
	}
	return -1
}

func findLowestCommonAncestor(node *TreeNode, p int, q int) *TreeNode {
	if node == nil || node.Val == p || node.Val == q {
		return node
	} else {
		var left = findLowestCommonAncestor(node.Left, p, q)
		var right = findLowestCommonAncestor(node.Right, p, q)
		if left != nil && right != nil {
			return node
		} else if left != nil {
			return left
		} else {
			return right
		}
	}
}

func findDistance(root *TreeNode, p int, q int) int {
	if p == q {
		return 0
	}
	var commonAncestor = findLowestCommonAncestor(root, p, q)
	var left = distanceTo(commonAncestor, p)
	var right = distanceTo(commonAncestor, q)
	return left + right
}
