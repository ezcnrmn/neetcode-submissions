type MinStack struct {
	stack []int
	min   *int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if ms.min == nil || *ms.min > val {
		ms.min = &val
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) >0 {
		poped := ms.stack[len(ms.stack)-1]
		ms.stack = ms.stack[:len(ms.stack)-1]

		if *ms.min == poped {
			ms.min = nil	
			for i := range ms.stack {
				if ms.min == nil || *ms.min > ms.stack[i] {
					ms.min = &ms.stack[i]
				}
			}
		}
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) > 0 {
		return ms.stack[len(ms.stack)-1]
	}
	return 0
}

func (ms *MinStack) GetMin() int {
	if ms.min != nil {
		return *ms.min
	}

	return 0
}
