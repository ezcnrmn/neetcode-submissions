class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        graph = defaultdict(list)
        for source, target, time in times:
            graph[source].append((time, target))
        
        dest_time = dict()
        dest_time[k] = 0
        min_heap = [(0, k)]
        result = 0
        while min_heap:
            time, node = heapq.heappop(min_heap)
            if dest_time.get(node, float('inf')) < time:
                continue
            result = max(result, time)
            for ngh_time, ngh in graph[node]:
                if dest_time.get(ngh, float('inf')) < ngh_time + time: 
                    continue
                
                heapq.heappush(min_heap, (ngh_time+time, ngh))
                dest_time[ngh] = ngh_time + time
        
        return result if len(dest_time) == n else -1