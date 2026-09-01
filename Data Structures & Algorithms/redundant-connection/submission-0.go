func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    graph := make(map[int][]int)
    indexes := make(map[[2]int]int)
    for i, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
        indexes[[2]int{edge[0], edge[1]}] = i
        indexes[[2]int{edge[1], edge[0]}] = i
    }

    // 0 - white, 1 - gray, 2 - black
    colors := make(map[int]int8)
    var resultIndex int

    var dfs func(int, int) bool
    dfs = func(node, from int) bool {
        colors[node] = 1
        for _, ngh := range graph[node] {
            if colors[ngh] == 2 || ngh == from {
                continue
            }
            if colors[ngh] == 1 || dfs(ngh, node) {
                i := indexes[[2]int{node, ngh}]
                resultIndex = max(i, resultIndex)
                return true
            } 
        }
        colors[node] = 2
        return false
    }

    for node := 1; node <= n; node++ {
        if colors[node] != 0 {
            continue
        }

        if dfs(node, -1) {
            break
        }
    }

    return edges[resultIndex]
}
