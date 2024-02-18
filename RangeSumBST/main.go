package RangeSumBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var thisValue int
	if root == nil {
		return thisValue
	}
	if root.Val >= low && root.Val <= high {
		thisValue = root.Val
	}
	if low < root.Val {
		thisValue += rangeSumBST(root.Left, low, high)
	}
	if high > root.Val {
		thisValue += rangeSumBST(root.Right, low, high)
	}
	return thisValue
}
