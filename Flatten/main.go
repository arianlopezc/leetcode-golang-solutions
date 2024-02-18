package Flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flattenAux(node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		return node
	}
	if node.Left != nil && node.Right != nil {
		var resLeft = flattenAux(node.Left)
		node.Right, resLeft.Right, node.Left = node.Left, node.Right, nil
		return flattenAux(node.Right)
	} else if node.Right != nil {
		return flattenAux(node.Right)
	} else {
		node.Right, node.Left = node.Left, nil
		return flattenAux(node.Right)
	}
}

func flatten(root *TreeNode) {
	if root != nil {
		flattenAux(root)
	}
}
