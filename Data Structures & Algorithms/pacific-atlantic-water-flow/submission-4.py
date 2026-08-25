class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        def dfs(row, col, seen):
            seen.add((row, col))
            for d in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
                newRow, newCol = row + d[0], col + d[1]
                if (
                    newRow < 0
                    or newRow >= len(heights)
                    or newCol < 0
                    or newCol >= len(heights[0])
                    or heights[newRow][newCol] < heights[row][col]
                    or (newRow, newCol) in seen
                ):
                    continue

                dfs(newRow, newCol, seen)

        pacific = set()
        atlantic = set()
        for i in range(len(heights[0])):
            dfs(0, i, pacific)
            dfs(len(heights)-1, i, atlantic)

        for i in range(0, len(heights)):
            dfs(i, 0, pacific)
            dfs(i, len(heights[0])-1, atlantic)

        result = []

        for i in range(len(heights)):
            row = []
            for j in range(len(heights[0])):
                el = (i, j)
                if el in pacific and el in atlantic:
                    row.append(("p", "a"))
                elif el in pacific:
                    row.append(("p", "."))
                elif el in atlantic:
                    row.append((".", "a"))
                else:
                    row.append((".", "."))
            print(row)

        for p in pacific & atlantic:
            result.append([p[0], p[1]])

        return result
