func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make(map[int][]int, numCourses)
    for _, p := range prerequisites {
        graph[p[1]] = append(graph[p[1]], p[0])
    }

    colors := make(map[int]int8, numCourses)

    for n := range numCourses {
        if colors[n] == 2 {
            continue
        }

        var dfs func(int) bool
        dfs = func(node int) bool {
            colors[node] = 1
            for _, neighbor := range graph[node] {
                color := colors[neighbor]
                if (color == 0 && !dfs(neighbor)) || color == 1{
                    return false
                } 
            }
            colors[node] = 2
            return true
        }
        if !dfs(n) {
            return false
        }
    }

    return true
}