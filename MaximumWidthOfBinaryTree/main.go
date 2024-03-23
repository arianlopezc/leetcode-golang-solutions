package MaximumWidthOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	var maxWidth int
	var children = []*TreeNode{root}
	for len(children) > 0 {
		var temp []*TreeNode
		var firstFilled bool
		var lastFilled int
		for _, child := range children {
			if child != nil {
				if !firstFilled && (child.Left != nil || child.Right != nil) {
					if child.Left != nil {
						temp = append(temp, child.Left)
						lastFilled = len(temp) - 1
					}
					temp = append(temp, child.Right)
					if child.Right != nil {
						lastFilled = len(temp) - 1
					}
					firstFilled = true
				} else if firstFilled {
					temp = append(temp, child.Left, child.Right)
					firstFilled = true
					if child.Right != nil {
						lastFilled = len(temp) - 1
					} else if child.Left != nil {
						lastFilled = len(temp) - 2
					}
				}
			} else if firstFilled {
				temp = append(temp, nil, nil)
			}
		}
		children = temp
		if maxWidth < lastFilled {
			maxWidth = lastFilled
		}
	}
	return maxWidth + 1
}
