type Edge struct {
    Node   int
    Weight int
}

type MinHeap []Edge

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].Weight < h[j].Weight
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make(map[int][]Edge, n)
    for _, time := range times {
        graph[time[0]] = append(graph[time[0]], Edge{Node: time[1], Weight: time[2]})
    }

    minHeap := &MinHeap{Edge{Node: k, Weight: 0}}
    seen := make(map[int]struct{}, n)
    var result int

    for minHeap.Len() > 0 {
        edge := (heap.Pop(minHeap)).(Edge)
        if _, visited := seen[edge.Node]; visited {
            continue
        }

        result = edge.Weight
        seen[edge.Node] = struct{}{}

        for _, nghEdge := range graph[edge.Node] {
            if _, visited := seen[nghEdge.Node]; !visited {
                heap.Push(minHeap, Edge{Node: nghEdge.Node, Weight: nghEdge.Weight + edge.Weight})
            }
        }
    }

    if len(seen) != n {
        return -1
    }

    return result
}
