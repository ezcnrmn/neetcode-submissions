type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type Queue struct {
	Head *Node
	Tail *Node
	Len  int
}

func newQueue() *Queue {
	head := &Node{}
	tail := &Node{Prev: head}
	head.Next = tail
	return &Queue{Head: head, Tail: tail}
}

func (q *Queue) Enqueue(x int) {
	node := &Node{Val: x}
	node.Next = q.Tail
	node.Prev = q.Tail.Prev
	q.Tail.Prev.Next = node
	q.Tail.Prev = node
	q.Len++
}

func (q *Queue) Dequeue() int {
	node := q.Head.Next
	q.Head.Next = node.Next
	node.Next.Prev = q.Head
	q.Len--
	return node.Val
}

type MyStack struct {
	Q *Queue
}

func Constructor() MyStack {
	return MyStack{Q: newQueue()}
}

func (this *MyStack) Push(x int) {
	this.Q.Enqueue(x)
}

func (this *MyStack) Pop() int {
	for range this.Q.Len - 1 {
		this.Q.Enqueue(this.Q.Dequeue())
	}
	return this.Q.Dequeue()
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.Q.Tail.Prev.Val
}

func (this *MyStack) Empty() bool {
	return this.Q.Len == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
