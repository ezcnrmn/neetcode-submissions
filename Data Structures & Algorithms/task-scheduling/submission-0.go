func leastInterval(tasks []byte, n int) int {
	taskTimes := make([]int, len(tasks))
	mem := make(map[byte]int)
	for i, t := range tasks {
		taskTimes[i] = mem[t]
		mem[t] += n+1
	}
	heap := New(taskTimes)
	var result int
	for len(heap.data) > 0 {
		time := heap.Pop()
		if time <= result {
			result++
		} else {
			result = time + 1
		}
	}
	return result
}

type MinHeap struct {
	data []int
}

func New(initial []int) *MinHeap {
	data := make([]int, len(initial))
	copy(data, initial)
	for i := (len(data) / 2) - 1; i >= 0; i-- {
		curI := i
		for curI*2+1 < len(data) {
			minChildI := curI*2 + 1
			if minChildI+1 < len(data) && data[minChildI+1] < data[minChildI] {
				minChildI++
			}
			if data[minChildI] < data[curI] {
				data[minChildI], data[curI] = data[curI], data[minChildI]
			} else {
				break
			}
			curI = minChildI
		}
	}
	return &MinHeap{data: data}
}

func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		return 0
	}

	poped := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	var i int
	for {
		if (i*2)+1 >= len(h.data) {
			break
		}

		minChildI := (i * 2) + 1
		if minChildI+1 < len(h.data) && h.data[minChildI+1] < h.data[minChildI] {
			minChildI++
		}

		if h.data[minChildI] < h.data[i] {
			h.data[minChildI], h.data[i] = h.data[i], h.data[minChildI]
		} else {
			break
		}
		i = minChildI
	}

	return poped
}

func (h *MinHeap) Push(value int) {
	h.data = append(h.data, value)
	i := len(h.data) - 1
	for {
		if i == 0 {
			break
		}
		parentI := (i - 1) / 2

		if h.data[parentI] < h.data[i] {
			break
		}

		h.data[parentI], h.data[i] = h.data[i], h.data[parentI]
		i = parentI
	}
}