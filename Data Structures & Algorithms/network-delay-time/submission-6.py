class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        graph = defaultdict(list)
        for source, target, time in times:
            graph[source].append((time, target))

        seen = set()
        min_heap = [(0, k)]
        result = 0
        while min_heap:
            time, node = heapq.heappop(min_heap)
            if node in seen:
                continue
            result = time
            seen.add(node)
            for ngh_time, ngh in graph[node]:
                if ngh not in seen:
                    heapq.heappush(min_heap, (ngh_time + time, ngh))
        
        if len(seen) != n:
            return -1
        return result
