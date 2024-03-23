package PathSumII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func paths(node *TreeNode, targetSum int, currSum int, path []int, res *[][]int) {
	if node.Left == nil && node.Right == nil {
		var pathTotal = currSum + node.Val
		if pathTotal == targetSum {
			*res = append(*res, append(path, node.Val))
		}
	} else {
		if node.Left != nil {
			paths(node.Left, targetSum, currSum+node.Val, append(path, node.Val), res)
		}
		if node.Right != nil {
			paths(node.Right, targetSum, currSum+node.Val, append(path, node.Val), res)
		}
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root != nil {
		paths(root, targetSum, 0, []int{}, &res)
	}
	return res
}
