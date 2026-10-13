type Node struct {
	Val  int
	Next *Node
	Min  *MinNode
}

type MinNode struct {
	Val  int
	Next *MinNode
}

type MinStack struct {
	Head *Node
	Min  *MinNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	node := &Node{Val: val, Next: ms.Head}
	ms.Head = node

	if ms.Min == nil || node.Val < ms.Min.Val {
		minNode := &MinNode{Val: val, Next: ms.Min}
		ms.Min = minNode
		node.Min = minNode
	}
}

func (ms *MinStack) Pop() {
	if ms.Head == nil {
		return
	}

	if ms.Head.Min != nil {
		ms.Min = ms.Min.Next
	}

	ms.Head = ms.Head.Next
}

func (ms *MinStack) Top() int {
	if ms.Head == nil {
		return 0
	}

	return ms.Head.Val
}

func (ms *MinStack) GetMin() int {
	if ms.Min == nil {
		return 0
	}

	return ms.Min.Val
}
