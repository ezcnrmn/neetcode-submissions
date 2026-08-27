func canFinish(numCourses int, prerequisites [][]int) bool {
    indegree := make(map[int]int, numCourses)
    graph := make(map[int][]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
        indegree[p[0]]++
        indegree[p[1]] = indegree[p[1]]
    }

    for {
        var queue []int
        for node, amount := range indegree {
            if amount == 0 {
                queue = append(queue, node)
                delete(indegree, node)
            }
        }

        if len(queue) == 0 {
            break
        }

        for _, node := range queue {
            for _, neighbor := range graph[node] {
                indegree[neighbor]--
            }
        }
    }

    return len(indegree) == 0 
}
