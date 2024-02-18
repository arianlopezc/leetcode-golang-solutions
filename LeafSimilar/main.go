package LeafSimilar

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeLeafs(node *TreeNode, leafs []int) []int {
	if node.Left == nil && node.Right == nil {
		leafs = append(leafs, node.Val)
		return leafs
	}
	if node.Left != nil {
		leafs = treeLeafs(node.Left, leafs)
	}
	if node.Right != nil {
		leafs = treeLeafs(node.Right, leafs)
	}
	return leafs
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil || root1 != nil && root2 == nil {
		return false
	}
	var leafsOne []int = treeLeafs(root1, []int{})
	var leafsTwo []int = treeLeafs(root2, []int{})
	if len(leafsTwo) != len(leafsOne) {
		return false
	}
	var index int = 0
	for index < len(leafsTwo) {
		if leafsOne[index] != leafsTwo[index] {
			return false
		}
		index++
	}
	return true
}
