func canFinish(numCourses int, prerequisites [][]int) bool {
    indegree := make(map[int]int)
    graph := make(map[int][]int, numCourses)
    for _, p := range prerequisites {
        from, to := p[1], p[0]
        graph[from] = append(graph[from], to)
        indegree[to]++
    }

    stack := make([]int, 0, numCourses)
    for course := range numCourses {
        if indegree[course] == 0 {
            stack = append(stack, course)
        }
    }

    var visited int
    for len(stack) > 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        visited++

        for _, nbr := range graph[cur] {
            indegree[nbr]--
            if indegree[nbr] == 0 {
                stack = append(stack, nbr)
            }
        }
    }

    return visited == numCourses
}
