package FindLargestValueInEachTreeRow

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int
	if root != nil {
		var children = []*TreeNode{root}
		for len(children) > 0 {
			var currMax = children[0].Val
			var temp []*TreeNode
			for _, child := range children {
				if child.Left != nil {
					temp = append(temp, child.Left)
				}
				if child.Right != nil {
					temp = append(temp, child.Right)
				}
				if child.Val > currMax {
					currMax = child.Val
				}
			}
			res = append(res, currMax)
			children = temp
		}
	}
	return res
}
