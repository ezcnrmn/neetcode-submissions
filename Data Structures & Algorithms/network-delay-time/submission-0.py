class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        graph = defaultdict(list)
        for source, target, time in times:
            graph[source].append((time, target))
        
        dest_time = dict()
        dest_time[k] = 0
        min_heap = [(0, k)]
        while min_heap:
            time, node = heapq.heappop(min_heap)
            if node in dest_time and dest_time[node] < time:
                continue
            
            for ngh_time, ngh in graph[node]:
                if ngh in dest_time and dest_time[ngh] < ngh_time + time: 
                    continue
                
                heapq.heappush(min_heap, (ngh_time+time, ngh))
                dest_time[ngh] = ngh_time + time
        
        result = 0
        for node in range(1, n+1):
            if node not in dest_time:
                return -1
            
            result = max(result, dest_time[node])
        
        return result