type Node struct{
    Num  int
    Time int
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Time < h[j].Time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make(map[int][]Node, n)
    for _, edge := range times {
        from, to, time := edge[0], edge[1], edge[2]
        graph[from] = append(graph[from], Node{Num: to, Time: time})
    }

    visited := make(map[int]struct{}, n)
    var result int
    h := &MinHeap{Node{Num: k, Time: 0}}
    for h.Len() > 0 && len(visited) < n {
        cur := heap.Pop(h).(Node)
        if _, seen := visited[cur.Num]; seen {
            continue
        }

        visited[cur.Num] = struct{}{}
        result = max(cur.Time, result)

        for _, nbr := range graph[cur.Num] {
            if _, seen := visited[nbr.Num]; !seen {
                heap.Push(h, Node{Num: nbr.Num, Time: cur.Time + nbr.Time})
            }
        }
    }

    if len(visited) != n {
        return -1
    }
    return result
}
