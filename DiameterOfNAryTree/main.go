package DiameterOfNAryTree

type Node struct {
	Val      int
	Children []*Node
}

func calculate(node *Node, maxD *int) int {
	if node != nil {
		var currMax int
		for _, child := range node.Children {
			var childMax = calculate(child, maxD)
			if currMax+childMax > *maxD {
				*maxD = currMax + childMax
			}
			if currMax < childMax {
				currMax = childMax
			}
		}
		return currMax + 1
	}
	return 0
}

func diameter(root *Node) int {
	var maxD int
	calculate(root, &maxD)
	return maxD
}
