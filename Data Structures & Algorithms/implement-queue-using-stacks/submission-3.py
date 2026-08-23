class MyQueue:
    def __init__(self):
        # append and pop only
        self.stack1 = collections.deque()
        self.stack2 = collections.deque()

    def push(self, x: int) -> None:
        self.stack1.append(x)

    def pop(self) -> int:
        if len(self.stack2) > 0:
            return self.stack2.pop()
        
        for _ in range(len(self.stack1)-1):
            self.stack2.append(self.stack1.pop())
        return self.stack1.pop()

    def peek(self) -> int:
        if len(self.stack2) > 0:
            return self.stack2[-1]
        if len(self.stack1) > 0:
            return self.stack1[0]
        return 0

    def empty(self) -> bool:
        return len(self.stack1) + len(self.stack2) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()