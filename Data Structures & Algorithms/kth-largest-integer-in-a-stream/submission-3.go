type KthLargest struct {
	Heap *MinHeap
	K    int
}


func Constructor(k int, nums []int) KthLargest {
	heap := NewMinHeap(nums)
	for len(heap.Data) > k {
		heap.Pop()
	}
	return KthLargest{
		Heap: heap,
		K:    k,
	}
}


func (this *KthLargest) Add(val int) int {
	this.Heap.Push(val)
	if len(this.Heap.Data) > this.K {
		this.Heap.Pop()
	}

	return this.Heap.Peek()
}

type MinHeap struct {
	Data []int
}

func NewMinHeap(initial []int) *MinHeap {
	data := make([]int, len(initial))
	copy(data, initial)
	for i := (len(data) / 2) - 1; i >= 0; i-- {
		curI := i
		for curI*2+1 < len(data) {
			maxChildI := curI*2 + 1
			if maxChildI+1 < len(data) && data[maxChildI+1] < data[maxChildI] {
				maxChildI++
			}
			if data[maxChildI] < data[curI] {
				data[maxChildI], data[curI] = data[curI], data[maxChildI]
			} else {
				break
			}
			curI = maxChildI
		}
	}
	return &MinHeap{Data: data}
}

func (h *MinHeap) Pop() int {
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

		minChildI := (i * 2) + 1
		if (i*2)+2 < len(h.Data) && h.Data[i*2+2] < h.Data[minChildI] {
			minChildI = (i * 2) + 2
		}

		if h.Data[minChildI] < h.Data[i] {
			h.Data[minChildI], h.Data[i] = h.Data[i], h.Data[minChildI]
		} else {
			break
		}
		i = minChildI
	}

	return poped
}

func (h *MinHeap) Push(value int) {
	h.Data = append(h.Data, value)
	i := len(h.Data) - 1
	for {
		if i == 0 {
			break
		}
		parentI := (i - 1) / 2

		if h.Data[parentI] < h.Data[i] {
			break
		}

		h.Data[parentI], h.Data[i] = h.Data[i], h.Data[parentI]
		i = parentI
	}
}

func (h *MinHeap) Peek() int {
	if len(h.Data) == 0 {
		return 0
	}
	return h.Data[0]
}