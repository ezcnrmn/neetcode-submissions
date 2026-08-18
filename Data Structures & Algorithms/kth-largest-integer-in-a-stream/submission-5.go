type KthLargest struct {
	Heap *MinHeap
	K    int
}


func Constructor(k int, nums []int) KthLargest {
	heap := NewMinHeap(nums)
	for len(heap.Data) > k {
		heap.Pop()
	}
	return KthLargest{Heap: heap, K: k}
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
	for i := len(data)/2-1; i >= 0; i-- {
		curI := i
		for {
			minChildI := curI*2+1
			if minChildI >= len(data) {
				break
			}
			if minChildI+1 < len(data) && data[minChildI+1] < data[minChildI] {
				minChildI++
			}
			if data[minChildI] < data[curI] {
				data[minChildI], data[curI] = data[curI], data[minChildI]
				curI = minChildI
			} else {
				break
			}
		}
	}
	return &MinHeap{Data: data}
}

func (mh *MinHeap) Push(value int) {
	mh.Data = append(mh.Data, value)
	curI := len(mh.Data)-1
	for {
		parentI := (curI-1)/2
		if parentI >= curI {
			break
		}
		if mh.Data[parentI] < mh.Data[curI] {
			break
		}
		mh.Data[parentI], mh.Data[curI] = mh.Data[curI], mh.Data[parentI]
		curI = parentI
	}
}

func (mh *MinHeap) Pop() int {
	if len(mh.Data) == 0 {
		return 0
	}

	poped := mh.Data[0]
	mh.Data[0] = mh.Data[len(mh.Data)-1]
	mh.Data = mh.Data[:len(mh.Data)-1]
	var curI int
	for {
		minChildI := curI*2+1
		if minChildI >= len(mh.Data) {
			break
		}
		if minChildI+1 < len(mh.Data) && mh.Data[minChildI+1] < mh.Data[minChildI] {
			minChildI++
		}
		if mh.Data[minChildI] < mh.Data[curI] {
			mh.Data[minChildI], mh.Data[curI] = mh.Data[curI], mh.Data[minChildI]
			curI = minChildI
		} else {
			break
		}
	}

	return poped
}

func (mh *MinHeap) Peek() int {
	if len(mh.Data) == 0 {
		return 0
	}
	return mh.Data[0]
}