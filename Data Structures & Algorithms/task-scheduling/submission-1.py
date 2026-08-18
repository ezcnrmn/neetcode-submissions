class Solution:
	def leastInterval(self, tasks: List[str], n: int) -> int:
		mem = dict()
		heap = []

		for t in tasks:
			time = mem.get(t, 0)
			heap.append(time)
			mem[t] = mem.get(t, 0)+n+1

		heapq.heapify(heap)
		counter = 0
		while len(heap) > 0:
			time = heapq.heappop(heap)
			if time > counter:
				counter = time+1
			else:
				counter += 1
		
		return counter
		