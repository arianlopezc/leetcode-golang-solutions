package BalanceABinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(node *TreeNode, order *[]int) {
	if node != nil {
		inOrderTraversal(node.Left, order)
		*order = append(*order, node.Val)
		inOrderTraversal(node.Right, order)
	}
}

func balanceArrayToBinaryTree(order []int) *TreeNode {
	if order == nil || len(order) == 0 {
		return nil
	}
	var center = len(order) / 2
	var node = &TreeNode{
		Val:   order[len(order)/2],
		Left:  balanceArrayToBinaryTree(order[:center]),
		Right: balanceArrayToBinaryTree(order[center+1:]),
	}
	return node
}

func balanceBST(root *TreeNode) *TreeNode {
	var order = make([]int, 0)
	inOrderTraversal(root, &order)
	return balanceArrayToBinaryTree(order)
}
