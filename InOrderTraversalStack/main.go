package InOrderTraversalStack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value       *TreeNode
		prev        *node
		VisitedLeft bool
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
	n := &node{value, this.top, false}
	this.top = n
	this.length++
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	var stack Stack
	if root != nil {
		stack.Push(root)
	}
	for stack.length > 0 {
		var topOfStack = stack.Peek()
		if topOfStack.Left != nil && !stack.top.VisitedLeft {
			stack.top.VisitedLeft = true
			stack.Push(topOfStack.Left)
		} else {
			result = append(result, topOfStack.Val)
			stack.Pop()
			if topOfStack.Right != nil {
				stack.Push(topOfStack.Right)
			}
		}
	}
	return result
}
