class Solution:
    def solve(self, board: List[List[str]]) -> None:
        safe = set()
        queue = collections.deque()
        for i in range(len(board[0])):
            p1, p2 = (0, i), (len(board) - 1, i)
            if board[p1[0]][p1[1]] == "O":
                queue.append(p1)
                safe.add(p1)
            if board[p2[0]][p2[1]] == "O":
                queue.append(p2)
                safe.add(p2)

        for i in range(1, len(board) - 1):
            p1, p2 = (i, 0), (i, len(board[0]) - 1)
            if board[p1[0]][p1[1]] == "O":
                queue.append(p1)
                safe.add(p1)
            if board[p2[0]][p2[1]] == "O":
                queue.append(p2)
                safe.add(p2)

        while len(queue) > 0:
            cur = queue.popleft()

            for d in [(1, 0), (0, 1), (-1, 0), (0, -1)]:
                newRow, newCol = cur[0] + d[0], cur[1] + d[1]
                if (
                    newRow < 0
                    or newRow >= len(board)
                    or newCol < 0
                    or newCol >= len(board[0])
                    or board[newRow][newCol] == "X"
                    or (newRow, newCol) in safe
                ):
                    continue

                safe.add((newRow, newCol))
                queue.append((newRow, newCol))

        for row in range(len(board)):
            for col in range(len(board[0])):
                if board[row][col] == "O" and (row, col) not in safe:
                    board[row][col] = "X"
