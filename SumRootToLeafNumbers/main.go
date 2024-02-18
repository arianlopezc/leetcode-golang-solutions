package SumRootToLeafNumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sum(node *TreeNode, total *int, subTotal int) {
	if node.Left == nil && node.Right == nil {
		var max = subTotal*10 + node.Val
		*total = max + *total
	} else {
		if node.Left != nil {
			sum(node.Left, total, subTotal*10+node.Val)
		}
		if node.Right != nil {
			sum(node.Right, total, subTotal*10+node.Val)
		}
	}
}

func sumNumbers(root *TreeNode) int {
	var total = 0
	sum(root, &total, 0)
	return total
}
