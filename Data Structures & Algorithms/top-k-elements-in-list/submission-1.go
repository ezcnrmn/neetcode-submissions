type Element struct{
	Num    int
	Amount int
}

type Heap []Element

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Amount < h[j].Amount }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any) {
	*h = append(*h, x.(Element))
}
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num]++
	}

	temp := make(Heap, 0, len(counter))
	for num, amount := range counter {
		temp = append(temp, Element{Num: num, Amount: -amount})
	}
	h := &temp

	heap.Init(h)

	result := make([]int, 0, k)
	for range k {
		el := (heap.Pop(h)).(Element)
		result = append(result, el.Num)
	}

	return result
}
