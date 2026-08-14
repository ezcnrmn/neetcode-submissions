func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap(stones)
	for len(heap.Data) > 1 {
		s1, s2 := heap.Pop(), heap.Pop()
		if s1 > s2 {
			heap.Push(s1-s2)
		} else if s2 > s1 {
			heap.Push(s2-s1)
		}
	}

	if len(heap.Data) == 1 {
		return heap.Peek()
	}

	return 0
}

type MaxHeap struct {
	Data []int
}

func NewMaxHeap(initial []int) *MaxHeap {
	data := make([]int, len(initial))
	copy(data, initial)
	for i := (len(data) / 2) - 1; i >= 0; i-- {
		curI := i
		for curI*2+1 < len(data) {
			maxChildI := curI*2 + 1
			if maxChildI+1 < len(data) && data[maxChildI+1] > data[maxChildI] {
				maxChildI++
			}
			if data[maxChildI] > data[curI] {
				data[maxChildI], data[curI] = data[curI], data[maxChildI]
			} else {
				break
			}
			curI = maxChildI
		}
	}
	return &MaxHeap{Data: data}
}

func (h *MaxHeap) Pop() int {
	if len(h.Data) == 0 {
		return 0
	}

	poped := h.Data[0]
	h.Data[0] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	var i int
	for {
		if (i*2)+1 >= len(h.Data) {
			break
		}

		maxChildI := (i * 2) + 1
		if maxChildI+1 < len(h.Data) && h.Data[maxChildI+1] > h.Data[maxChildI] {
			maxChildI++
		}

		if h.Data[maxChildI] > h.Data[i] {
			h.Data[maxChildI], h.Data[i] = h.Data[i], h.Data[maxChildI]
		} else {
			break
		}
		i = maxChildI
	}

	return poped
}

func (h *MaxHeap) Push(value int) {
	h.Data = append(h.Data, value)
	i := len(h.Data) - 1
	for {
		if i == 0 {
			break
		}
		parentI := (i - 1) / 2

		if h.Data[parentI] > h.Data[i] {
			break
		}

		h.Data[parentI], h.Data[i] = h.Data[i], h.Data[parentI]
		i = parentI
	}
}

func (h *MaxHeap) Peek() int {
	if len(h.Data) == 0 {
		return 0
	}
	return h.Data[0]
}