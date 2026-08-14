type KthLargest struct {
	Heap *MaxHeap
	K    int
}


func Constructor(k int, nums []int) KthLargest {
	heap := &MaxHeap{}
	for _, number := range nums {
		heap.Push(number)
	}
	return KthLargest{
		Heap: heap,
		K:    k,
	}
}


func (this *KthLargest) Add(val int) int {
	this.Heap.Push(val)

	poped := make([]int, 0, this.K-1)
	for range this.K-1 {
		poped = append(poped, this.Heap.Pop())
	}
	result := this.Heap.Peek()
	for len(poped) > 0 {
		cur := poped[len(poped)-1]
		poped = poped[:len(poped)-1]
		this.Heap.Push(cur)
	}
	return result
}

type MaxHeap struct {
	Data []int
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
		if (i*2)+2 < len(h.Data) && h.Data[i*2+2] > h.Data[maxChildI] {
			maxChildI = (i * 2) + 2
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
