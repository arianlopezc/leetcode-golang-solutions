package NaryTreePreorderTraversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	var stack []*Node
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		var node = stack[0]
		res = append(res, node.Val)
		stack = stack[1:]
		for index := len(node.Children) - 1; index >= 0; index-- {
			stack = append([]*Node{node.Children[index]}, stack...)
		}
	}
	return res
}
