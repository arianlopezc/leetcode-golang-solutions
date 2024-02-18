package LongestZigZag

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigZag(node *TreeNode, left bool, subPath int) int {
	if left && node.Left != nil {
		return zigZag(node.Left, false, subPath+1)
	} else if !left && node.Right != nil {
		return zigZag(node.Right, true, subPath+1)
	} else {
		return subPath
	}
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var totals []int
	totals = append(totals, zigZag(root, true, 0))
	totals = append(totals, zigZag(root, false, 0))
	totals = append(totals, longestZigZag(root.Left))
	totals = append(totals, longestZigZag(root.Right))
	var max int
	for _, val := range totals {
		if max < val {
			max = val
		}
	}
	return max
}
