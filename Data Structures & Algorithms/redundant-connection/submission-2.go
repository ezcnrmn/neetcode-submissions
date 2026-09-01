func findRedundantConnection(edges [][]int) []int {
    n := len(edges)

    graph := make(map[int][]int, n)
    indexes := make(map[[2]int]int, n*2)
    indegree := make(map[int]int, n)
    for i, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
        indexes[[2]int{edge[0], edge[1]}] = i
        indexes[[2]int{edge[1], edge[0]}] = i
        indegree[edge[0]]++
        indegree[edge[1]]++
    }

    queue := make([]int, 0, n)
    for node, i := range indegree {
        if i == 1 {
            queue = append(queue, node)
            delete(indegree, node)
        }
    }

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        for _, ngh := range graph[node] {
            if _, ok := indegree[ngh]; !ok {
                continue
            }
            indegree[ngh]--
            if indegree[ngh] == 1 {
                queue = append(queue, ngh)
                delete(indegree, ngh)
            }
        }
    }

    var resultIndex int
    for node := range indegree {
        for _, ngh := range graph[node] {
            if _, ok := indegree[ngh]; ok {
                resultIndex = max(resultIndex, indexes[[2]int{node, ngh}])
            }
        }
    }

    return edges[resultIndex]
}