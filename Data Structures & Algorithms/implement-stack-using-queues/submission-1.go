type MyStack struct {
	queue q
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue.enqueue(x)
	for i := 0; i < len(this.queue.s)-1; i++ {
		this.queue.enqueue(this.queue.dequeue())
	}
}

func (this *MyStack) Pop() int {
	return this.queue.dequeue()
}

func (this *MyStack) Top() int {
	return this.queue.s[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue.s) == 0
}

// queue for MyStack

type q struct {
	s []int
}

func (q *q) enqueue(x int) {
	q.s	= append(q.s, x)
}

func (q *q) dequeue() int {
	v := q.s[0]
	q.s = q.s[1:]
	return v
}
