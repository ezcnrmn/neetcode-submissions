func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make(map[int]int, numCourses)
	graph := make(map[int][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indegree[p[0]]++
	}

	var queue []int
	for course := range numCourses {
		if indegree[course] == 0 {
			queue = append(queue, course)
		}
	}

	result := make([]int, 0, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		result = append(result, course)
		for _, neighbor := range graph[course] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != numCourses {
		return []int{}
	}

	return result
}
