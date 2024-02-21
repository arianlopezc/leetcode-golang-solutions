package BinaryTreePostOrderTraversalStack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value        *TreeNode
		prev         *node
		VisitedLeft  bool
		VisitedRight bool
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() *TreeNode {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() *TreeNode {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value *TreeNode) {
	n := &node{value, this.top, false, false}
	this.top = n
	this.length++
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	var stack Stack
	if root != nil {
		stack.Push(root)
	}
	for stack.length > 0 {
		var topOfStack = stack.top
		if topOfStack.value.Right != nil && !topOfStack.VisitedRight {
			stack.Push(topOfStack.value.Right)
		}
		if topOfStack.value.Left != nil && !topOfStack.VisitedLeft {
			stack.Push(topOfStack.value.Left)
		}
		if !topOfStack.VisitedRight && !topOfStack.VisitedLeft {
			topOfStack.VisitedRight = true
			topOfStack.VisitedLeft = true
			continue
		}
		result = append(result, topOfStack.value.Val)
		stack.Pop()
	}
	return result
}
