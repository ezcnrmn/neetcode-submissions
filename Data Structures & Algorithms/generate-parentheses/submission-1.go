type qItem struct {
	str    string
	opened int
	closed int
}

func generateParenthesis(n int) []string {
	queue := []qItem{{"", 0, 0}}
	for queue[0].opened < n || queue[0].closed < n {
		cur := queue[0]
		queue = queue[1:]

		if cur.opened < n {
			queue = append(queue, qItem{str: cur.str + "(", opened: cur.opened+1, closed: cur.closed})
		}
		if cur.closed < n && cur.opened > cur.closed {
			queue = append(queue, qItem{str: cur.str + ")", opened: cur.opened, closed: cur.closed+1})
		}
	}

	result := make([]string, 0, len(queue))
	for _, item := range queue {
		result = append(result, item.str)
	}

	return result
}
