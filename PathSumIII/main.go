package PathSumIII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	return getPathSums(root, []int{}, targetSum)
}

func getPathSums(n *TreeNode, path []int, target int) int {
	if n == nil {
		return 0
	}
	np := append(path, n.Val)
	temp := 0
	sum := 0
	for i := len(np) - 1; i >= 0; i-- {
		sum += np[i]
		if sum == target {
			temp++
		}
	}
	return getPathSums(n.Left, np, target) + getPathSums(n.Right, np, target) + temp
}
