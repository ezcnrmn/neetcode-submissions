class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        graph = defaultdict(list)
        for from_node, to_node in edges:
            graph[from_node].append(to_node)
            graph[to_node].append(from_node)

        stack = collections.deque()
        seen = set()
        result = 0
        for node in range(n):
            if node in seen:
                continue

            result += 1
            stack.append(node)
            while stack:
                cur = stack.pop()
                for neighbor in graph[cur]:
                    if neighbor in seen:
                        continue
                    seen.add(neighbor)
                    stack.append(neighbor)

        return result
