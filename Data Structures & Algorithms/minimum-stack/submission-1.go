type Min struct {
	val int
	pos int
}

type MinStack struct {
	s    []int
	mins []Min
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	m.s = append(m.s, val)
	if len(m.mins) == 0 || m.mins[len(m.mins)-1].val > val {
		m.mins = append(m.mins, Min{val: val, pos: len(m.s)-1})
	}
}

func (m *MinStack) Pop() {
	if m.mins[len(m.mins)-1].pos == len(m.s)-1 {
		m.mins = m.mins[:len(m.mins)-1]
	}
	m.s = m.s[:len(m.s)-1]
}

func (m *MinStack) Top() int {
	return m.s[len(m.s)-1]
}

func (m *MinStack) GetMin() int {
	return m.mins[len(m.mins)-1].val
}
