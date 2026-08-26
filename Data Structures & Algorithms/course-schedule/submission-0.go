func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make(map[int][]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
    }

    current := make(map[int]struct{}, numCourses)
    visited := make(map[int]struct{}, numCourses)

    for n := range numCourses {
        if _, seen := visited[n]; seen {
            continue
        }

        stack := []int{n}
        for len(stack) > 0 {
            cur := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if _, seen := current[cur]; seen {
                return false
            }
            current[cur] = struct{}{}
            stack = append(stack, graph[cur]...)
        }

        for node := range current {
            delete(current, node)
            visited[node] = struct{}{}
        } 
    }

    return true
}