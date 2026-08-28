class Solution:
	def validTree(self, n: int, edges: List[List[int]]) -> bool:
		if n-1 != len(edges):
			return False

		indegree = defaultdict(lambda: 0)
		graph = defaultdict(list)
		for from_node, to_node in edges:
			graph[from_node].append(to_node)
			indegree[to_node] += 1

		queue = collections.deque()
		for node in range(n):
			if indegree[node] > 1:
				return False
			if indegree[node] == 0:
				queue.append(node)
				del indegree[node]
		
		while queue:
			node = queue.popleft()
			for neighbor in graph[node]:
				indegree[neighbor] -= 1
				queue.append(neighbor)
				del indegree[neighbor]
		
		if len(indegree) > 0:
			return False
		
		return True