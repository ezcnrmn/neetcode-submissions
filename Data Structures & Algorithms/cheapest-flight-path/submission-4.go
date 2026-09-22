func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    graph := make(map[int][][2]int, n)
    for _, flight := range flights {
        graph[flight[0]] = append(graph[flight[0]], [2]int{flight[1], flight[2]})
    }

    dists := make([]int, 0, n)
    for range n {
        dists = append(dists, math.MaxInt)
    }
    dists[src] = 0

    minHeap := &MinHeap{{Num: src}}
    result := math.MaxInt

    for minHeap.Len() > 0{
        cur := heap.Pop(minHeap).(HeapItem)
        
        for _, flight := range graph[cur.Num] {
            flightDst, flightPrice := flight[0], flight[1]

            if dists[flightDst] < cur.Price + flightPrice {
                continue
            }

            if dst == flightDst {
                result = min(result, cur.Price + flightPrice)
                continue
            }

            if cur.Jumps + 1 > k  {
                continue
            }

            heap.Push(minHeap, HeapItem{Num: flightDst, Price: cur.Price + flightPrice, Jumps: cur.Jumps+1})
        }
    }

    if result == math.MaxInt {
        return -1
    }
    return result
}

/*
h := &MinHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
fmt.Printf("minimum: %d\n", (*h)[0])
for h.Len() > 0 {
    fmt.Printf("%d ", heap.Pop(h))
}
*/


type HeapItem struct {
    Num   int
    Price int
    Jumps int
}
type MinHeap [] HeapItem

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].Price < h[j].Price
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(HeapItem))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
