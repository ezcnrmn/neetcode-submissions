type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

type MyQueue struct {
	A stack
	B stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.A.push(x)
}

func (this *MyQueue) Pop() int {
	if len(this.B) == 0 {
		for len(this.A) > 0 {
			this.B.push(this.A.pop())
		}
	}

	return this.B.pop()
}

func (this *MyQueue) Peek() int {
	if len(this.B) == 0 {
		for len(this.A) > 0 {
			this.B.push(this.A.pop())
		}
	}

	return this.B[len(this.B)-1]
}

func (this *MyQueue) Empty() bool {
	return 0 == len(this.A) + len(this.B)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
