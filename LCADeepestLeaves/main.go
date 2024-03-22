package LCADeepestLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lca(node *TreeNode, depth int, maxDepth *int) *TreeNode {
	if node.Left == nil && node.Right == nil && depth == *maxDepth {
		if depth == *maxDepth {
			return node
		}
		return nil
	} else {
		var left *TreeNode
		if node.Left != nil {
			left = lca(node.Left, depth+1, maxDepth)
		}
		var right *TreeNode
		if node.Right != nil {
			right = lca(node.Right, depth+1, maxDepth)
		}
		if right != nil && left != nil {
			return node
		} else if right != nil {
			return right
		} else {
			return left
		}
	}
}

func depth(node *TreeNode) int {
	var children = []*TreeNode{node}
	var maxDepth = 1
	for len(children) > 0 {
		var temp []*TreeNode
		for _, child := range children {
			if child.Right != nil {
				temp = append(temp, child.Right)
			}
			if child.Left != nil {
				temp = append(temp, child.Left)
			}
		}
		if len(temp) > 0 {
			maxDepth++
		}
		children = temp
	}
	return maxDepth
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var maxDepth = depth(root)
	return lca(root, 1, &maxDepth)
}
