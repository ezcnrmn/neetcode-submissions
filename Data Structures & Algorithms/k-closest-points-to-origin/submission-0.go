func kClosest(points [][]int, k int) [][]int {
	pointsByDistance:= make(map[float64][][]int)
	for _, p := range points {
		x, y := float64(p[0]), float64(p[1])
		distance := math.Sqrt(x*x+y*y)
		pointsByDistance[distance] = append(pointsByDistance[distance], p)
	}
	heap := &MinHeap{}
	for distance := range pointsByDistance {
		heap.Push(distance)
	}
	result := make([][]int, 0, k)
	var i int
	for i < k{
		distance := heap.Pop()
		result = append(result, pointsByDistance[distance]...)
		i += len(pointsByDistance[distance])
	}
	return result
}

type MinHeap struct {
	data []float64
}

func (h *MinHeap) Pop() float64 {
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

func (h *MinHeap) Push(value float64) {
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
