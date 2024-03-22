package IsCompleteTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	var children = []*TreeNode{root}
	var noMore = false
	for len(children) > 0 {
		var temp []*TreeNode
		for _, child := range children {
			if child.Right != nil && child.Left == nil {
				return false
			} else {
				if !noMore && ((child.Right == nil && child.Left == nil) || (child.Right == nil && child.Left != nil)) {
					noMore = true
				} else {
					if noMore && (child.Left != nil || child.Right != nil) {
						return false
					}
				}
			}
			if child.Left != nil {
				temp = append(temp, child.Left)
			}
			if child.Right != nil {
				temp = append(temp, child.Right)
			}
		}
		children = temp
	}
	return true
}
