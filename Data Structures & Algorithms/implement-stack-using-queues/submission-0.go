type MyStack struct {
	s []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.s = append(this.s, x)
}

func (this *MyStack) Pop() int {
	v := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	return v
}

func (this *MyStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.s) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
