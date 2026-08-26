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
            for _, neighbor := range graph[cur] {
                if _, seen := visited[cur]; seen {
                    continue
                }
                stack = append(stack, neighbor)
            }
        }

        for node := range current {
            delete(current, node)
            visited[node] = struct{}{}
        } 
    }

    return true
}