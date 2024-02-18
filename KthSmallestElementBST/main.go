package KthSmallestElementBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findKthSmallest(root *TreeNode, k int, val *int) int {
	if root == nil {
		if k == 0 {
			return root.Val
		} else {
			return k
		}
	} else {
		if root.Left != nil {
			k = findKthSmallest(root.Left, k, val)
			if k == 0 {
				return 0
			}
		}
		k--
		if k == 0 {
			*val = root.Val
			return 0
		}
		if root.Right != nil {
			k = findKthSmallest(root.Right, k, val)
			if k == 0 {
				return 0
			}
		}
		return k
	}
}

func kthSmallest(root *TreeNode, k int) int {
	var res int
	findKthSmallest(root, k, &res)
	return res
}
