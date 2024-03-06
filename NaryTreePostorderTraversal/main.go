package NaryTreePostorderTraversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	var stack []*Node
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		var node = stack[0]
		if node.Children != nil && len(node.Children) > 0 {
			for index := len(node.Children) - 1; index >= 0; index-- {
				stack = append([]*Node{node.Children[index]}, stack...)
			}
			node.Children = nil
		} else {
			res = append(res, node.Val)
			stack = stack[1:]
		}
	}
	return res
}
