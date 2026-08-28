class Solution:
	def validTree(self, n: int, edges: List[List[int]]) -> bool:

		graph = defaultdict(list)
		for from_node, to_node in edges:
			graph[from_node].append(to_node)
			graph[to_node].append(from_node)

		queue = collections.deque([(0, -1)])
		visited = set([0])
		
		while queue:
			node, from_node = queue.popleft()
			for neighbor in graph[node]:
				if neighbor == from_node:
					continue
				if neighbor in visited:
					return False
				queue.append((neighbor, node))
				visited.add(neighbor)

		if len(visited) != n:
			return False
		
		return True