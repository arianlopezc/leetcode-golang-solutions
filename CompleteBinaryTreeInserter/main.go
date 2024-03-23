package CompleteBinaryTreeInserter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	CompleteTree []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	var cbt = CBTInserter{}
	var children = []*TreeNode{root}
	for len(children) > 0 {
		var temp []*TreeNode
		for _, child := range children {
			cbt.CompleteTree = append(cbt.CompleteTree, child)
			if child.Left != nil {
				temp = append(temp, child.Left)
			}
			if child.Right != nil {
				temp = append(temp, child.Right)
			}
		}
		children = temp
	}
	return cbt
}

func (this *CBTInserter) Insert(val int) int {
	var newNode = &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	this.CompleteTree = append(this.CompleteTree, newNode)
	if len(this.CompleteTree) == 1 {
		return 0
	}
	var pos = len(this.CompleteTree) - 1
	var parent = (pos - 1) / 2
	if this.CompleteTree[parent].Left == nil {
		this.CompleteTree[parent].Left = newNode
	} else {
		this.CompleteTree[parent].Right = newNode
	}
	return this.CompleteTree[parent].Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	if len(this.CompleteTree) > 0 {
		return this.CompleteTree[0]
	}
	return nil
}
