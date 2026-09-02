class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        parent = [i for i in range(n)]
        rank = [1 for _ in range(n)]

        def find(node):
            cur = node
            while cur != parent[cur]:
                parent[cur] = parent[parent[cur]]
                cur = parent[cur]
            return cur

        def union(node1, node2):
            par1, par2 = find(node1), find(node2)
            if par1 == par2:
                return False  # not unified

            if rank[par1] > rank[par2]:
                parent[par2] = par1
                rank[par1] += rank[par2]
            else:
                parent[par1] = par2
                rank[par2] += rank[par1]
            return True  # unified

        for node1, node2 in edges:
            if union(node1, node2):
                n -= 1

        return n
