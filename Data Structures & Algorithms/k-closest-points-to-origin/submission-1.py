class Solution:
	def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
		mem = dict()
		heap = []
		for p in points:
			distance = math.sqrt(p[0]*p[0] + p[1]*p[1])
			if distance not in heap:
				heapq.heappush(heap, distance)
			if distance in mem:
				mem[distance].append(p)
			else:
				mem[distance] = [p]
		
		result = []
		while len(result) < k:
			distance = heapq.heappop(heap)
			for p in mem[distance]:
				result.append(p)
			
		return result
