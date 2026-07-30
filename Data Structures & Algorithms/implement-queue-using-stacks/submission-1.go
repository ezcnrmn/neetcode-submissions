type MyQueue struct {
	stack s
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	temp := s{}
	n := len(this.stack.s)
	for i := 0; i < n; i++ {
		temp.push(this.stack.pop())
	}
	this.stack.push(x)
	for i := 0; i < n; i++ {
		this.stack.push(temp.pop())
	}
}

func (this *MyQueue) Pop() int {
	return this.stack.pop()
}

func (this *MyQueue) Peek() int {
	return this.stack.s[len(this.stack.s)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack.s) == 0
}

// stack for MyQueue

type s struct {
	s []int
}

func (s *s) push(x int) {
	s.s = append(s.s, x)
}

func (s *s) pop() int {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}
