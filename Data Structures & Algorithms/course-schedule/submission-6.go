func canFinish(numCourses int, prerequisites [][]int) bool {
    indegree := make(map[int]int, numCourses)
    graph := make(map[int][]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
        indegree[p[0]]++
        indegree[p[1]] = indegree[p[1]]
    }

    var queue []int
    for node, amount := range indegree {
        if amount == 0 {
            queue = append(queue, node)
            delete(indegree, node)
        }
    }

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, neighbor := range graph[cur] {
            indegree[neighbor]--
            if indegree[neighbor] == 0 {
                queue = append(queue, neighbor)
                delete(indegree, neighbor)
            }
        }
    }

    return len(indegree) == 0 
}
