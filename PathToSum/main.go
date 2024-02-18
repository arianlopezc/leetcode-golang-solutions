package PathToSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func nav(root *TreeNode, targetSum int, pathSum int) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val+pathSum == targetSum
	} else {
		var leftResult bool
		if root.Left != nil {
			leftResult = nav(root.Left, targetSum, pathSum+root.Val)
		}
		if leftResult {
			return true
		}
		if root.Right != nil {
			return nav(root.Right, targetSum, pathSum+root.Val)
		}
		return false
	}
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return nav(root, targetSum, 0)
}
