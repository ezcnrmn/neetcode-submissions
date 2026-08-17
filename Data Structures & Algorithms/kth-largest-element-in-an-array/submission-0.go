func findKthLargest(nums []int, k int) int {
	heap := New(nums)
	for range k-1 {
		heap.Pop()
	}
	return heap.data[0]
}
type Heap struct {
	data []int
}

func New(initial []int) *Heap {
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
	return &Heap{data: data}
}

func (h *Heap) Pop() int {
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

		maxChildI := (i * 2) + 1
		if maxChildI+1 < len(h.data) && h.data[maxChildI+1] > h.data[maxChildI] {
			maxChildI++
		}

		if h.data[maxChildI] > h.data[i] {
			h.data[maxChildI], h.data[i] = h.data[i], h.data[maxChildI]
		} else {
			break
		}
		i = maxChildI
	}

	return poped
}
