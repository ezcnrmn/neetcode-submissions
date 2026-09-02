class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        parent = [i for i in range(n)]
        rank = [1 for _ in range(n)]

        def find(point):
            if parent[point] != point:
                parent[point] = find(parent[point]) 
            return parent[point]
        
        def union(point1, point2):
            par1, par2 = find(point1), find(point2)
            if par1 == par2:
                return False
            
            if rank[par1] > rank[par2]:
                parent[par2] = par1
                rank[par1] += rank[par2]
            else:
                parent[par1] = par2
                rank[par2] += rank[par1]
            return True
        
        def manhattan(point1, point2):
            return abs(points[point1][0] - points[point2][0]) + abs(points[point1][1] - points[point2][1])
        
        distances = []
        for point1 in range(n):
            for point2 in range(point1 + 1, n):
                distances.append((manhattan(point1, point2), point1, point2))
        heapq.heapify(distances)

        result = 0
        while distances:
            distance, point1, point2 = heapq.heappop(distances)
            if union(point1, point2):
                result += distance
        
        return result
